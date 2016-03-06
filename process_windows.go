package kiwi

// TODO: switch from
// github.com/Andoryuuta/w32 to github.com/VividCortex/w32 if pull request is accepted.

import (
	"errors"
	"fmt"
	"github.com/Andoryuuta/w32"
	"path/filepath"
	"reflect"
	"unsafe"
)

// Platform specific fields to be embeded into
// the Process struct.
type ProcPlatAttribs struct {
	Handle w32.HANDLE
}

// Constant for full process access
const PROCESS_ALL_ACCESS = w32.PROCESS_VM_READ | w32.PROCESS_VM_WRITE | w32.PROCESS_VM_OPERATION | w32.PROCESS_QUERY_INFORMATION

//getFileNameByPID returns a file name given a PID.
func getFileNameByPID(pid uint32) string {
	var fileName string = `<Unknown File>`

	//Open process
	hnd, ok := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, pid)
	if !ok {
		return fileName
	}
	defer w32.CloseHandle(hnd)

	//Get file path
	path, ok := w32.GetProcessImageFileName(hnd)
	if !ok {
		return fileName
	}

	//Split file path to get file name
	_, fileName = filepath.Split(path)
	return fileName
}

// GetProcessByFileName returns the process with the given file name.
// If multiple processes have the same filename, the first process
// enumerated by this function is returned.
func GetProcessByFileName(fileName string) (Process, error) {
	// Read in process ids
	PIDs := make([]uint32, 1024)
	var bytesRead uint32 = 0
	ok := w32.EnumProcesses(PIDs, uint32(len(PIDs)), &bytesRead)
	if !ok {
		panic("Error Enumerating processes.")
	}

	// Loop over pids,
	// Divide bytesRead by sizeof(uint32) to get how many processes there are.
	for i := uint32(0); i < (bytesRead / 4); i++ {
		// Make sure to skip over the system process with PID 0
		if PIDs[i] == 0 {
			continue
		}

		// Check if its the process
		if getFileNameByPID(PIDs[i]) == fileName {
			hnd, ok := w32.OpenProcess(PROCESS_ALL_ACCESS, false, PIDs[i])
			if !ok {
				return Process{}, errors.New(fmt.Sprintf("Error while opening process %d", PIDs[i]))
			}
			return Process{hnd, PIDs[i]}, nil
		}
	}

	// Couldn't find process, return an error
	return Process{}, errors.New("Couldn't find process with name " + fileName)
}

// Taken from genkman's gist(https://gist.github.com/henkman/3083408)
func (p *Process) GetModuleBase(moduleName string) (uintptr, error) {
	snap := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE32|w32.TH32CS_SNAPALL|w32.TH32CS_SNAPMODULE, p.PID)
	if snap == 0 {
		return 0, errors.New("Error trying on create toolhelp32 snapshot.")
	}
	defer w32.CloseHandle(snap)

	var me32 w32.MODULEENTRY32
	me32.Size = uint32(unsafe.Sizeof(me32))

	// Get first module
	if !w32.Module32First(snap, &me32) {
		return 0, errors.New("Error trying to get first module.")
	}

	// Check first module
	if syscall.UTF16ToString(me32.SzModule[:]) == moduleName {
		return uintptr(unsafe.Pointer(me32.ModBaseAddr)), nil
	}

	// Loop all modules remaining
	for w32.Module32Next(snap, &me32) {
		// Check this module
		if syscall.UTF16ToString(me32.SzModule[:]) == moduleName {
			return uintptr(unsafe.Pointer(me32.ModBaseAddr)), nil
		}
	}

	// If this is reached, then we couldn't find the module
	return 0, errors.New("Couldn't Find Module.")
}

// The Windows specific read function.
func (p *Process) read(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	i := reflect.Indirect(v)
	size := i.Type().Size()
	bytesRead, ok := w32.ReadProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(i.UnsafeAddr()),
		size,
	)
	if !ok || bytesRead != size {
		return errors.New("Error on reading process memory.")
	}
	return nil
}

// The Windows specific write function.
func (p *Process) write(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	i := reflect.Indirect(v)
	size := i.Type().Size()
	bytesWritten, ok := w32.WriteProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(i.UnsafeAddr()),
		size,
	)
	if !ok || bytesWritten != size {
		return errors.New("Error on writing process memory.")
	}
	return nil
}
