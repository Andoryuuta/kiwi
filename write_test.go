package kiwi

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name string

		runTest func(p Process, t *testing.T) (error, interface{}, interface{})
	}{
		{
			name: "uint8",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected uint8 = 243
				var outVar uint8 = 0
				err := p.WriteUint8(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "uint16",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected uint16 = 65523
				var outVar uint16 = 0
				err := p.WriteUint16(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "uint32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected uint32 = 4282681632
				var outVar uint32 = 0
				err := p.WriteUint32(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "uint64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected uint64 = 18437214073702121416
				var outVar uint64 = 0
				err := p.WriteUint64(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "int8",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected int8 = 125
				var outVar int8 = 0
				err := p.WriteInt8(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "int16",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected int16 = 32524
				var outVar int16 = 0
				err := p.WriteInt16(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "int32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected int32 = 2121625523
				var outVar int32 = 0
				err := p.WriteInt32(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "int32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected int32 = 2121625523
				var outVar int32 = 0
				err := p.WriteInt32(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "int64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected int64 = 9217263856192656271
				var outVar int64 = 0
				err := p.WriteInt64(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "float32",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected float32 = 2515123.0321
				var outVar float32 = 0
				err := p.WriteFloat32(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "float64",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				var expected float64 = 782658917272.1512
				var outVar float64 = 0
				err := p.WriteFloat64(uintptr(unsafe.Pointer(&outVar)), expected)
				return err, outVar, expected
			},
		},
		{
			name: "bytes",
			runTest: func(p Process, t *testing.T) (error, interface{}, interface{}) {
				expected := []byte{5, 4, 3, 2, 1}
				outVar := make([]byte, 5)
				err := p.WriteBytes(uintptr(unsafe.Pointer(&outVar[0])), expected)
				return err, outVar, expected
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
				t.Fatalf("Error trying to write. Error: %s\n", err.Error())
			}

			// Check the results.
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("Written value does not match expected. Got: %v, Expected: %v\n", got, want)
			}
		})
	}
}
