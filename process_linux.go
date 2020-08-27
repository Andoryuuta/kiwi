package kiwi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	_ "path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Platform specific fields to be embedded into
// the Process struct.
type ProcPlatAttribs struct {
}

// GetProcessByPID returns the process with the given PID.
func GetProcessByPID(PID int) (Process, error) {
	// Try to open the folder to see if it exists and if we have access to it.
	f, err := os.Open(fmt.Sprintf("/proc/%d", PID))
	defer f.Close()
	if err != nil {
		return Process{}, errors.New(fmt.Sprintf("Error opening /proc/%d", PID))
	}

	return Process{PID: uint64(PID)}, nil
}

// GetProcessByFileName returns the process with the given file name.
// If multiple processes have the same filename, the first process
// enumerated by this function is returned.
func GetProcessByFileName(fileName string) (Process, error) {
	// Open the /proc *nix directory.
	procDir, err := os.Open("/proc")
	defer procDir.Close()
	if err != nil {
		fmt.Println(err)
		return Process{}, errors.New("Error on opening /proc")
	}

	// Read in the directory names. Processes are listed
	// as folders within this directory, named by their PID.
	dirNames, err := procDir.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return Process{}, errors.New("Error on reading dirs from /proc")
	}

	// Enumerate all directories here.
	for _, dirString := range dirNames {

		// Parse the directory name as a uint.
		pid, err := strconv.ParseUint(dirString, 0, 64)

		// If it is not a numeric directory name, skip it.
		if v, ok := err.(*strconv.NumError); ok && v.Err == strconv.ErrSyntax {
			continue
		} else if err != nil {
			fmt.Println(err)
			return Process{}, errors.New("Error on enumerating dirs from /proc")
		}

		// TODO: Change this to something better, it is very hacky right now.

		// Read the target program's stats file.
		tmpFileBytes, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/stat", pid))
		if err != nil {
			fmt.Println(err)
			return Process{}, errors.New("Error on enumerating dirs from " + fmt.Sprintf("/proc/%d/stat", pid))
		}

		// HACK!
		// Stat file [1] has the file name surrounded by ()
		curProcFileName := strings.Trim(strings.Split(string(tmpFileBytes), " ")[1], "()")

		// Check if this is the process we are looking for.
		if curProcFileName == fileName {
			return Process{PID: pid}, nil
		}
	}

	return Process{}, errors.New("Couldn't find a process with the given file name.")
}

// The platform specific read function.
func (p *Process) read(addr uintptr, ptr interface{}) error {
	// Reflection magic!
	v := reflect.ValueOf(ptr)
	dataAddr := getDataAddr(v)
	dataSize := getDataSize(v)

	// Open the file mapped process memory.
	mem, err := os.Open(fmt.Sprintf("/proc/%d/mem", p.PID))
	defer mem.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening /proc/%d/mem. Are you root?", p.PID))
	}

	// Create a buffer and read data into it.
	dataBuf := make([]byte, dataSize)
	n, err := mem.ReadAt(dataBuf, int64(addr))
	if n != int(dataSize) {
		return errors.New(fmt.Sprintf("Tried to read %d bytes, actually read %d bytes\n", dataSize, n))
	} else if err != nil {
		return err
	}

	// Unsafely cast to []byte to copy data into.
	buf := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: dataAddr,
		Len:  int(dataSize),
		Cap:  int(dataSize),
	}))
	copy(*buf, dataBuf)
	return nil
}

// The platform specific write function.
func (p *Process) write(addr uintptr, ptr interface{}) error {
	// Reflection magic!
	v := reflect.ValueOf(ptr)
	dataAddr := getDataAddr(v)
	dataSize := getDataSize(v)

	// Open the file mapped process memory.
	mem, err := os.OpenFile(fmt.Sprintf("/proc/%d/mem", p.PID), os.O_WRONLY, 0666)
	defer mem.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening /proc/%d/mem. Are you root?", p.PID))
	}

	// Unsafe cast to []byte to copy data from.
	buf := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: dataAddr,
		Len:  int(dataSize),
		Cap:  int(dataSize),
	}))

	// Write the data from buf into memory.
	n, err := mem.WriteAt(*buf, int64(addr))
	if n != int(dataSize) {
		return errors.New((fmt.Sprintf("Tried to write %d bytes, actually wrote %d bytes\n", dataSize, n)))
	} else if err != nil {
		return err
	}
	return nil
}
