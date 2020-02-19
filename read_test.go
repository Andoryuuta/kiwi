package kiwi

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/kardianos/osext"
)

// The currently running process's name,
// Set in TestMain.
var currentProcessName string

func TestMain(m *testing.M) {
	// Get current executable name.
	fn, err := osext.Executable()
	if err != nil {
		panic(fmt.Sprintf("Error trying to get executable name. Error: %s\n", err.Error()))
	}
	currentProcessName = filepath.Base(fn)

	// Run tests and exit.
	os.Exit(m.Run())
}

func TestGetProcessByPID(t *testing.T) {
	// Get process using kiwi.
	pid := os.Getpid()
	_, err := GetProcessByPID(pid)
	if err != nil {
		t.Fatalf("Error trying to open process with PID %d, Error: %s\n", pid, err.Error())
	}
}

func TestGetProcessByFileName(t *testing.T) {
	// Get process using kiwi.
	_, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}
}

func TestReadUint8(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar uint8 = 243

	// Attempt to read using kiwi.
	readVar, err := p.ReadUint8(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadUint16(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar uint16 = 65523

	// Attempt to read using kiwi.
	readVar, err := p.ReadUint16(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadUint32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar uint32 = 4282681632

	// Attempt to read using kiwi.
	readVar, err := p.ReadUint32(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadUint64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar uint64 = 18437214073702121416

	// Attempt to read using kiwi.
	readVar, err := p.ReadUint64(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadInt8(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar int8 = 125

	// Attempt to read using kiwi.
	readVar, err := p.ReadInt8(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadInt16(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar int16 = 32524

	// Attempt to read using kiwi.
	readVar, err := p.ReadInt16(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadInt32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar int32 = 2121625523

	// Attempt to read using kiwi.
	readVar, err := p.ReadInt32(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadInt64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar int64 = 9217263856192656271

	// Attempt to read using kiwi.
	readVar, err := p.ReadInt64(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadFloat32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	var orgVar float32 = 2515123.0321

	// Attempt to read using kiwi.
	readVar, err := p.ReadFloat32(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadFloat64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	orgVar := 782658917272.1512

	// Attempt to read using kiwi.
	readVar, err := p.ReadFloat64(uintptr(unsafe.Pointer(&orgVar)))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if readVar != orgVar {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}

func TestReadBytes(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to read from.
	orgVar := []byte{5, 4, 3, 2, 1}

	// Attempt to read using kiwi.
	readVar, err := p.ReadBytes(uintptr(unsafe.Pointer(&orgVar[0])), len(orgVar))
	if err != nil {
		t.Fatalf("Error trying to read. Error: %s\n", err.Error())
	}

	if !bytes.Equal(readVar, orgVar) {
		t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", orgVar, readVar)
	}
}
