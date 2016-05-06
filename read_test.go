package kiwi

import (
	"fmt"
	"github.com/kardianos/osext"
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

// The currently running process's name,
// Set in TestMain.
var currentProcessName string

func TestMain(m *testing.M) {
	// Get curent executable name
	fn, err := osext.Executable()
	if err != nil {
		panic(fmt.Sprintf("Error trying to get executable name. Error: %s\n", err.Error()))
	}
	currentProcessName = filepath.Base(fn)

	// Run tests and exit
	os.Exit(m.Run())
}

func TestGetProcessByFileName(t *testing.T) {
	// Get process using kiwi
	_, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}
}

func TestReadUint8(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var uint8 = 243

	// Attempt to read using kiwi.
	read_var, err := p.ReadUint8(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadUint16(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var uint16 = 65523

	// Attempt to read using kiwi.
	read_var, err := p.ReadUint16(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadUint32(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var uint32 = 4282681632

	// Attempt to read using kiwi.
	read_var, err := p.ReadUint32(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadUint64(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var uint64 = 18437214073702121416

	// Attempt to read using kiwi.
	read_var, err := p.ReadUint64(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadInt8(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var int8 = 125

	// Attempt to read using kiwi.
	read_var, err := p.ReadInt8(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadInt16(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var int16 = 32524

	// Attempt to read using kiwi.
	read_var, err := p.ReadInt16(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadInt32(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var int32 = 2121625523

	// Attempt to read using kiwi.
	read_var, err := p.ReadInt32(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadInt64(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var int64 = 9217263856192656271

	// Attempt to read using kiwi.
	read_var, err := p.ReadInt64(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadFloat32(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var float32 = 2515123.0321

	// Attempt to read using kiwi.
	read_var, err := p.ReadFloat32(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadFloat64(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var float64 = 782658917272.1512

	// Attempt to read using kiwi.
	read_var, err := p.ReadFloat64(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}

func TestReadUintptr(t *testing.T) {
	// Get process using kiwi
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var org_var uintptr = 0x41414141

	// Attempt to read using kiwi.
	read_var, err := p.ReadUintptr(uintptr(unsafe.Pointer(&org_var)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if read_var != org_var {
		t.Fatalf("Read values are not the same. Orginial: %v, Read: %v\n", org_var, read_var)
	}
}
