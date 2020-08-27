package kiwi

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/Andoryuuta/kiwi/w32"
	"golang.org/x/sys/windows"
)

// ProcPlatAttribs contains platform specific fields to be
// embedded into the Process struct.
type ProcPlatAttribs struct {
	Handle w32.HANDLE
}

// NeededProcessAccess is the combined win32 process open flags needed for kiwi functionality.
const NeededProcessAccess = w32.PROCESS_VM_READ | w32.PROCESS_VM_WRITE | w32.PROCESS_VM_OPERATION | w32.PROCESS_QUERY_INFORMATION

// GetProcessByPID returns the process with the given PID.
func GetProcessByPID(pid int) (Process, error) {
	hnd, ok := w32.OpenProcess(NeededProcessAccess, false, uint32(pid))
	if !ok {
		return Process{}, fmt.Errorf("OpenProcess %v: %w", pid, windows.GetLastError())
	}
	return Process{ProcPlatAttribs: ProcPlatAttribs{Handle: hnd}, PID: uint64(pid)}, nil
}

// getFileNameByPID returns a file name given a PID.
func getFileNameByPID(pid uint32) (string, error) {
	// Open process.
	hnd, ok := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, pid)
	if !ok {
		return "", fmt.Errorf("OpenProcess %v: %w", pid, windows.GetLastError())
	}
	defer w32.CloseHandle(hnd)

	// Get file path.
	path, ok := w32.GetProcessImageFileName(hnd)
	if !ok {
		return "", fmt.Errorf("GetProcessImageFileName: %w", windows.GetLastError())
	}

	// Split file path to get file name.
	_, fileName := filepath.Split(path)
	return fileName, nil
}

// GetProcessByFileName returns the process with the given file name.
// If multiple processes have the same filename, the first process
// enumerated by this function is returned.
func GetProcessByFileName(fileName string) (Process, error) {
	// Read in process ids
	pidCount := 1024
	var pids []uint32
	var bytesRead uint32

	uint32size := uint32(unsafe.Sizeof(uint32(0)))

	// Get the process ids, increasing the pids buffer each time if there isn't enough space.
	for i := 1; uint32(len(pids))*uint32size == bytesRead; i++ {
		pids = make([]uint32, pidCount*i)
		ok := w32.EnumProcesses(pids, uint32(len(pids))*uint32size, &bytesRead)
		if !ok {
			return Process{}, fmt.Errorf("EnumProcesses: %w", windows.GetLastError())
		}
	}

	// Loop over pids,
	// (Divide bytesRead by sizeof(uint32) to get how many processes there are).
	for i := uint32(0); i < (bytesRead / uint32size); i++ {
		// Skip over the system process with PID 0.
		if pids[i] == 0 {
			continue
		}

		// Get the filename for this process
		curFileName, err := getFileNameByPID(pids[i])
		if err != nil {
			//return Process{}, fmt.Errorf("getFileNameByPID %v: %w", pids[i], err)
			continue
		}

		// Check if it is the process being searched for.
		if curFileName == fileName {
			hnd, ok := w32.OpenProcess(NeededProcessAccess, false, pids[i])
			if !ok {
				return Process{}, fmt.Errorf("OpenProcess %v: %w", pids[i], windows.GetLastError())
			}
			return Process{ProcPlatAttribs: ProcPlatAttribs{Handle: hnd}, PID: uint64(pids[i])}, nil
		}
	}

	// Couldn't find process, return an error.
	return Process{}, errors.New("couldn't find process with name " + fileName)
}

// GetModuleBase takes a module name as an argument. (e.g. "kernel32.dll")
// Returns the modules base address.
//
// (Mostly taken from genkman's gist: https://gist.github.com/henkman/3083408)
// TODO(Andoryuuta): Figure out possible licencing issues with this, or rewrite.
func (p *Process) GetModuleBase(moduleName string) (uintptr, error) {
	snap, ok := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE32|w32.TH32CS_SNAPALL|w32.TH32CS_SNAPMODULE, uint32(p.PID))
	if !ok {
		return 0, fmt.Errorf("CreateToolhelp32Snapshot: %w", windows.GetLastError())
	}
	defer w32.CloseHandle(snap)

	var me32 w32.MODULEENTRY32
	me32.DwSize = uint32(unsafe.Sizeof(me32))

	// Get first module.
	if !w32.Module32First(snap, &me32) {
		return 0, fmt.Errorf("Module32First: %w", windows.GetLastError())
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
	return 0, errors.New("couldn't find module")
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
		return errors.New("error reading process memory")
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
		return errors.New("error writing process memory")
	}
	return nil
}
