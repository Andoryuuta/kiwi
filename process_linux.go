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

// Platform specific fields to be embeded into
// the Process struct.
type ProcPlatAttribs struct {
}

// GetProcessByFileName returns the process with the given file name.
// If multiple processes have the same filename, the first process
// enumerated by this function is returned.
func GetProcessByFileName(fileName string) (Process, error) {
	// Open the /proc *nix directory.
	procDir, err := os.Open("/proc")
	if err != nil {
		fmt.Println(err)
		return Process{}, errors.New("Error on opening /proc")
	}

	// Read in the directory names, processes are listed
	// as folders here, named their PID's.
	dirNames, err := procDir.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return Process{}, errors.New("Error on reading dirs from /proc")
	}

	// Enumerate all directories here
	for _, dirString := range dirNames {

		// Parse the directory name as a uint
		pid, err := strconv.ParseUint(dirString, 0, 64)

		// If it is not a numeric dir name, skip it.
		if v, ok := err.(*strconv.NumError); ok && v.Err == strconv.ErrSyntax {
			continue
		} else if err != nil {
			fmt.Println(err)
			return Process{}, errors.New("Error on enumerating dirs from /proc")
		}

		// TODO: Change this to something better,
		// it is very hacky right now.

		// Read the target program's stats file
		tmpFileBytes, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/stat", pid))
		if err != nil {
			fmt.Println(err)
			return Process{}, errors.New("Error on enumerating dirs from " + fmt.Sprintf("/proc/%d/stat", pid))
		}

		// HACK!
		// Stat file [1] has the file name surrounded by ()
		curProcFileName := strings.Trim(strings.Split(string(tmpFileBytes), " ")[1], "()")

		//fmt.Printf("Pid: %d\tFile Name:%s\n", pid, curProcFileName)

		// Check if this is the process we are looking for.
		if curProcFileName == fileName {
			return Process{PID: pid}, nil
		}
	}

	return Process{}, errors.New("Couldn't find a process with the given file name.")
}

// The platform specific read function.
func (p *Process) read(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	i := reflect.Indirect(v)
	size := i.Type().Size()

	mem, err := os.Open(fmt.Sprintf("/proc/%d/mem", p.PID))
	if err != nil {
		// TODO: Return proper error
		panic(err)
	}

	// Create a buffer and read data into it
	dataBuf := make([]byte, size)
	n, err := mem.ReadAt(dataBuf, int64(addr))
	if n != int(size) {
		panic(fmt.Sprintf("Tried to read %d bytes, actually read %d bytes\n", size, n))
	} else if err != nil {
		panic(err)
	}

	buf := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: i.UnsafeAddr(),
		Len:  int(size),
		Cap:  int(size),
	}))
	copy(*buf, dataBuf)

	fmt.Println(buf)
	fmt.Println(dataBuf)

	return nil
}

// The platform specific write function.
func (p *Process) write(addr uintptr, ptr interface{}) error {
	panic("Not implemented")
	return nil
}
