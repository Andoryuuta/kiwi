package kiwi

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/Andoryuuta/kiwi/w32"
)

// Platform specific fields to be embedded into
// the Process struct.
type ProcPlatAttribs struct {
	Handle w32.HANDLE
}

// Constant for full process access.
const PROCESS_ALL_ACCESS = w32.PROCESS_VM_READ | w32.PROCESS_VM_WRITE | w32.PROCESS_VM_OPERATION | w32.PROCESS_QUERY_INFORMATION

// getFileNameByPID returns a file name given a PID.
func getFileNameByPID(pid uint32) string {
	var fileName string = `<Unknown File>`

	// Open process.
	hnd, ok := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, pid)
	if !ok {
		return fileName
	}
	defer w32.CloseHandle(hnd)

	// Get file path.
	path, ok := w32.GetProcessImageFileName(hnd)
	if !ok {
		return fileName
	}

	// Split file path to get file name.
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

	// Loop over PIDs,
	// Divide bytesRead by sizeof(uint32) to get how many processes there are.
	for i := uint32(0); i < (bytesRead / 4); i++ {
		// Skip over the system process with PID 0.
		if PIDs[i] == 0 {
			continue
		}

		// Check if it is the process being searched for.
		if getFileNameByPID(PIDs[i]) == fileName {
			hnd, ok := w32.OpenProcess(PROCESS_ALL_ACCESS, false, PIDs[i])
			if !ok {
				return Process{}, errors.New(fmt.Sprintf("Error while opening process %d", PIDs[i]))
			}
			return Process{ProcPlatAttribs: ProcPlatAttribs{Handle: hnd}, PID: uint64(PIDs[i])}, nil
		}
	}

	// Couldn't find process, return an error.
	return Process{}, errors.New("Couldn't find process with name " + fileName)
}

// GetModuleBase takes a module name as an argument. (e.g. "kernel32.dll")
// Returns the modules base address.
//
// (Mostly taken from genkman's gist: https://gist.github.com/henkman/3083408)
// TODO(Andoryuuta): Figure out possible licencing issues with this, or rewrite.
func (p *Process) GetModuleBase(moduleName string) (uintptr, error) {
	snap, ok := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE32|w32.TH32CS_SNAPALL|w32.TH32CS_SNAPMODULE, uint32(p.PID))
	if !ok {
		return 0, errors.New("Error trying on create toolhelp32 snapshot.")
	}
	defer w32.CloseHandle(snap)

	var me32 w32.MODULEENTRY32
	me32.DwSize = uint32(unsafe.Sizeof(me32))

	// Get first module.
	if !w32.Module32First(snap, &me32) {
		return 0, errors.New("Error trying to get first module.")
	}

	// Check first module.
	if syscall.UTF16ToString(me32.SzModule[:]) == moduleName {
		return uintptr(unsafe.Pointer(me32.ModBaseAddr)), nil
	}

	// Loop all modules remaining.
	for w32.Module32Next(snap, &me32) {
		// Check this module.
		if syscall.UTF16ToString(me32.SzModule[:]) == moduleName {
			return uintptr(unsafe.Pointer(me32.ModBaseAddr)), nil
		}
	}

	// Module couldn't be found.
	return 0, errors.New("Couldn't find module.")
}

// The platform specific read function.
func (p *Process) read(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	dataAddr := getDataAddr(v)
	dataSize := getDataSize(v)
	bytesRead, ok := w32.ReadProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(dataAddr),
		dataSize,
	)
	if !ok || bytesRead != dataSize {
		return errors.New("Error on reading process memory.")
	}
	return nil
}

// The platform specific write function.
func (p *Process) write(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	dataAddr := getDataAddr(v)
	dataSize := getDataSize(v)
	bytesWritten, ok := w32.WriteProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(dataAddr),
		dataSize,
	)
	if !ok || bytesWritten != dataSize {
		return errors.New("Error on writing process memory.")
	}
	return nil
}
