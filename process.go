package kiwi

import (
	"fmt"

	"golang.org/x/text/encoding/unicode"
)

// Process holds general information about the process,
// as well as embedding the struct ProcPlatAttribs which contains platform-specific data
// such as Windows process handles, linux `/proc` file handles, etc.
type Process struct {
	// Embedded struct for platform specific fields
	ProcPlatAttribs

	// Platform independent process details
	PID uint64
}

// ReadInt8 reads an int8.
func (p *Process) ReadInt8(addr uintptr) (int8, error) {
	var v int8
	e := p.read(addr, &v)
	return v, e
}

// ReadInt16 reads an int16.
func (p *Process) ReadInt16(addr uintptr) (int16, error) {
	var v int16
	e := p.read(addr, &v)
	return v, e
}

// ReadInt32 reads an int32.
func (p *Process) ReadInt32(addr uintptr) (int32, error) {
	var v int32
	e := p.read(addr, &v)
	return v, e
}

// ReadInt64 reads an int64
func (p *Process) ReadInt64(addr uintptr) (int64, error) {
	var v int64
	e := p.read(addr, &v)
	return v, e
}

// ReadUint8 reads an uint8.
func (p *Process) ReadUint8(addr uintptr) (uint8, error) {
	var v uint8
	e := p.read(addr, &v)
	return v, e
}

// ReadUint16 reads an uint16.
func (p *Process) ReadUint16(addr uintptr) (uint16, error) {
	var v uint16
	e := p.read(addr, &v)
	return v, e
}

// ReadUint32 reads an uint32.
func (p *Process) ReadUint32(addr uintptr) (uint32, error) {
	var v uint32
	e := p.read(addr, &v)
	return v, e
}

// ReadUint64 reads an uint64.
func (p *Process) ReadUint64(addr uintptr) (uint64, error) {
	var v uint64
	e := p.read(addr, &v)
	return v, e
}

// ReadFloat32 reads a float32.
func (p *Process) ReadFloat32(addr uintptr) (float32, error) {
	var v float32
	e := p.read(addr, &v)
	return v, e
}

// ReadFloat64 reads a float64
func (p *Process) ReadFloat64(addr uintptr) (float64, error) {
	var v float64
	e := p.read(addr, &v)
	return v, e
}

// ReadUint32Ptr reads a uint32 pointer chain with offsets.
func (p *Process) ReadUint32Ptr(addr uintptr, offsets ...uintptr) (uint32, error) {
	curPtr, err := p.ReadUint32(addr)
	if err != nil {
		return 0, fmt.Errorf("ReadUint32 0x%X: %w", addr, err)
	}

	for _, offset := range offsets {
		curPtr, err = p.ReadUint32(uintptr(curPtr) + offset)
		if err != nil {
			return 0, fmt.Errorf("ReadUint32 0x%X: %w", addr, err)
		}
	}

	return curPtr, nil
}

// ReadBytes reads a slice of bytes.
func (p *Process) ReadBytes(addr uintptr, size int) ([]byte, error) {
	v := make([]byte, size)
	e := p.read(addr, &v)
	return v, e
}

// Takes a []byte and returns the index of the first 0 byte, or -1 if none.
func cstrlen(data []byte) int {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			return i
		}
	}
	return -1
}

// ReadNullTerminatedUTF8String reads a null-termimated UTF8 string.
func (p *Process) ReadNullTerminatedUTF8String(addr uintptr) (string, error) {
	var outputBuffer []byte
	readSize := 2048
	for {
		// Make a buffer to hold our temporary string data.
		v := make([]byte, readSize)
		err := p.read(addr+uintptr(len(outputBuffer)), &v)
		if err != nil {
			// Halve readSize if we got an error on read.
			if readSize > 1 {
				readSize /= 2
				continue
			}

			return "", err
		}

		// Find the null index in our buffer.
		idx := cstrlen(v)
		if idx == -1 {
			// No null-terminator in the buffer.
			outputBuffer = append(outputBuffer, v...)
		} else {
			outputBuffer = v[:idx]
			break
		}
	}
	return string(outputBuffer), nil // Golang will decode as utf8.
}

// Takes a []uint16 and returns the index of the first 0, or -1 if none.
func cwstrlen(data []uint16) int {
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			return i
		}
	}
	return -1
}

// ReadNullTerminatedUTF16String reads a null-termimated UTF16 string.
// Respects BOM, assumes little endianess if no BOM is present.
func (p *Process) ReadNullTerminatedUTF16String(addr uintptr) (string, error) {
	var outputBuffer []uint16
	readSize := 2048
	for {
		// Make a buffer to hold our temporary string data.
		v := make([]uint16, readSize)
		err := p.read(addr+uintptr(len(outputBuffer)), &v)
		if err != nil {
			// Halve readSize if we got an error on read.
			if readSize > 1 {
				readSize /= 2
				continue
			}

			return "", err
		}

		// Find the null index in our buffer.
		idx := cwstrlen(v)
		if idx == -1 {
			// No null-terminator in the buffer.
			outputBuffer = append(outputBuffer, v...)
		} else {
			outputBuffer = v[:idx]
			break
		}
	}

	// Convert the []uint16 -> []byte
	byteBuf := make([]byte, len(outputBuffer)*2)
	for i := 0; i < len(outputBuffer); i++ {
		byteBuf[(i*2)+0] = byte((outputBuffer[i] >> 0) & 0xFF)
		byteBuf[(i*2)+1] = byte((outputBuffer[i] >> 8) & 0xFF)
	}

	// Decode the UTF16 to UTF8
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	utf8Bytes, err := decoder.Bytes(byteBuf)
	if err != nil {
		return "", err
	}

	return string(utf8Bytes), nil
}

// WriteInt8 writes an int8.
func (p *Process) WriteInt8(addr uintptr, v int8) error {
	return p.write(addr, &v)
}

// WriteInt16 writes an int16.
func (p *Process) WriteInt16(addr uintptr, v int16) error {
	return p.write(addr, &v)
}

// WriteInt32 writes an int32.
func (p *Process) WriteInt32(addr uintptr, v int32) error {
	return p.write(addr, &v)
}

// WriteInt64 writes an int64.
func (p *Process) WriteInt64(addr uintptr, v int64) error {
	return p.write(addr, &v)
}

// WriteUint8 writes an uint8.
func (p *Process) WriteUint8(addr uintptr, v uint8) error {
	return p.write(addr, &v)
}

// WriteUint16 writes an uint16.
func (p *Process) WriteUint16(addr uintptr, v uint16) error {
	return p.write(addr, &v)
}

// WriteUint32 writes an uint32.
func (p *Process) WriteUint32(addr uintptr, v uint32) error {
	return p.write(addr, &v)
}

// WriteUint64 writes an uint64.
func (p *Process) WriteUint64(addr uintptr, v uint64) error {
	return p.write(addr, &v)
}

// WriteFloat32 writes a float32.
func (p *Process) WriteFloat32(addr uintptr, v float32) error {
	return p.write(addr, &v)
}

// WriteFloat64 writes a float64.
func (p *Process) WriteFloat64(addr uintptr, v float64) error {
	return p.write(addr, &v)
}

// WriteBytes writes a slice of bytes.
func (p *Process) WriteBytes(addr uintptr, v []byte) error {
	return p.write(addr, &v)
}
