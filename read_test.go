package kiwi

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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

func TestRead(t *testing.T) {
	tests := []struct {
		name string

		runTest func(p Process, t *testing.T) (error, interface{}, interface{})
	}{
		{
			name: "uint8",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar uint8 = 243
				readVar, err := p.ReadUint8(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "uint16",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar uint16 = 65523
				readVar, err := p.ReadUint16(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "uint32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar uint32 = 4282681632
				readVar, err := p.ReadUint32(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "uint64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar uint64 = 18437214073702121416
				readVar, err := p.ReadUint64(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "int8",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar int8 = 125
				readVar, err := p.ReadInt8(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "int16",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar int16 = 32524
				readVar, err := p.ReadInt16(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "int32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar int32 = 2121625523
				readVar, err := p.ReadInt32(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "int64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar int64 = 9217263856192656271
				readVar, err := p.ReadInt64(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "float32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar float32 = 2515123.0321
				readVar, err := p.ReadFloat32(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "float64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var orgVar float64 = 782658917272.1512
				readVar, err := p.ReadFloat64(uintptr(unsafe.Pointer(&orgVar)))
				return err, orgVar, readVar
			},
		},
		{
			name: "bytes",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				orgVar := []byte{5, 4, 3, 2, 1}
				readVar, err := p.ReadBytes(uintptr(unsafe.Pointer(&orgVar[0])), len(orgVar))
				return err, orgVar, readVar
			},
		},
		{
			name: "null_terminated_utf8_string",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				expected := "Hello, 世界"
				orgVar := append([]byte(expected), 0x00)
				readVar, err := p.ReadNullTerminatedUTF8String(uintptr(unsafe.Pointer(&orgVar[0])))
				return err, readVar, expected
			},
		},
		{
			name: "null_terminated_utf16_string",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				expected := "0123"
				orgVar := []byte{0x30, 0x00, 0x31, 0x00, 0x32, 0x00, 0x33, 0x00, 0x00, 0x00}
				readVar, err := p.ReadNullTerminatedUTF16String(uintptr(unsafe.Pointer(&orgVar[0])))
				return err, readVar, expected
			},
		},
		{
			name: "null_terminated_utf16_string_bigendian_bom",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				expected := "0123"
				orgVar := []byte{0xFE, 0xFF, 0x00, 0x30, 0x00, 0x31, 0x00, 0x32, 0x00, 0x33, 0x00, 0x00}
				readVar, err := p.ReadNullTerminatedUTF16String(uintptr(unsafe.Pointer(&orgVar[0])))
				return err, readVar, expected
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Get process using kiwi.
			p, err := GetProcessByFileName(currentProcessName)
			if err != nil {
				t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
			}

			// Run the test.
			err, got, want := tst.runTest(p, t)
			if err != nil {
				t.Fatalf("Error trying to read. Error: %s\n", err.Error())
			}

			// Check the results.
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("Read values are not the same. Original: %v, Read: %v\n", got, want)
			}
		})
	}
}
