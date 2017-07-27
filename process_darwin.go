package kiwi

import (
	_ "errors"
	_ "fmt"
	_ "path/filepath"
	_ "reflect"
	_ "unsafe"
)

// Platform specific fields to be embedded into
// the Process struct.
type ProcPlatAttribs struct {
}

// GetProcessByPID returns the process with the given PID.
func GetProcessByPID(PID int) (Process, error) {
	panic("OSX is not supported")
	return Process{}, nil
}

// GetProcessByFileName returns the process with the given file name.
// If multiple processes have the same filename, the first process
// enumerated by this function is returned.
func GetProcessByFileName(fileName string) (Process, error) {
	panic("OSX is not supported")
	return Process{}, nil
}

// The platform specific read function.
func (p *Process) read(addr uintptr, ptr interface{}) error {
	panic("OSX is not supported")
	return nil
}

// The platform specific write function.
func (p *Process) write(addr uintptr, ptr interface{}) error {
	panic("OSX is not supported")
	return nil
}
