package kiwi

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestWriteUint8(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected uint8 = 243
	var outVar uint8 = 0

	// Attempt to read using kiwi.
	err = p.WriteUint8(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteUint16(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected uint16 = 65523
	var outVar uint16 = 0

	// Attempt to read using kiwi.
	err = p.WriteUint16(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteUint32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected uint32 = 4282681632
	var outVar uint32 = 0

	// Attempt to read using kiwi.
	err = p.WriteUint32(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteUint64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected uint64 = 18437214073702121416
	var outVar uint64 = 0

	// Attempt to read using kiwi.
	err = p.WriteUint64(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteInt8(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected int8 = 125
	var outVar int8 = 0

	// Attempt to read using kiwi.
	err = p.WriteInt8(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteInt16(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected int16 = 32524
	var outVar int16 = 0

	// Attempt to read using kiwi.
	err = p.WriteInt16(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteInt32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected int32 = 2121625523
	var outVar int32 = 0

	// Attempt to read using kiwi.
	err = p.WriteInt32(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteInt64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected int64 = 9217263856192656271
	var outVar int64 = 0

	// Attempt to read using kiwi.
	err = p.WriteInt64(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteFloat32(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected float32 = 2515123.0321
	var outVar float32 = 0

	// Attempt to read using kiwi.
	err = p.WriteFloat32(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteFloat64(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	var expected float64 = 782658917272.1512
	var outVar float64 = 0

	// Attempt to read using kiwi.
	err = p.WriteFloat64(uintptr(unsafe.Pointer(&outVar)), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if outVar != expected {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}

func TestWriteBytes(t *testing.T) {
	// Get process using kiwi.
	p, err := GetProcessByFileName(currentProcessName)
	if err != nil {
		t.Fatalf("Error trying to open process \"%s\", Error: %s\n", currentProcessName, err.Error())
	}

	// In memory variable to write to.
	expected := []byte{5, 4, 3, 2, 1}
	outVar := make([]byte, 5)

	// Attempt to read using kiwi.
	err = p.WriteBytes(uintptr(unsafe.Pointer(&outVar[0])), expected)
	if err != nil {
		t.Fatalf("Error trying to write. Error: %s\n", err.Error())
	}

	if bytes.Compare(outVar, expected) != 0 {
		t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", outVar, expected)
	}
}
