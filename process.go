package kiwi

import (
	"errors"
)

type Process struct {
	// Embedded struct for platform specific fields
	ProcPlatAttribs

	// Platform independent process details
	PID uint64
}

func (p *Process) ReadInt8(addr uintptr) (int8, error) {
	var v int8
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadInt16(addr uintptr) (int16, error) {
	var v int16
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadInt32(addr uintptr) (int32, error) {
	var v int32
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadInt64(addr uintptr) (int64, error) {
	var v int64
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadUint8(addr uintptr) (uint8, error) {
	var v uint8
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadUint16(addr uintptr) (uint16, error) {
	var v uint16
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadUint32(addr uintptr) (uint32, error) {
	var v uint32
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadUint64(addr uintptr) (uint64, error) {
	var v uint64
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadFloat32(addr uintptr) (float32, error) {
	var v float32
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadFloat64(addr uintptr) (float64, error) {
	var v float64
	e = p.read(addr, &v)
	return v, e
}

func (p *Process) ReadUint32Ptr(addr uintptr, offsets ...uintptr) (uint32, error) {
	curPtr, err := p.ReadUint32(addr)
	if err != nil {
		return 0, errors.New("Error while trying to read from ptr base.")
	}

	for _, offset := range offsets {
		curPtr, err = p.ReadUint32(uintptr(curPtr) + offset)
		if err != nil {
			return 0, errors.New("Error while trying to read from offset.")
		}
	}

	return curPtr, nil
}

func (p *Process) ReadBytes(addr uintptr, size int) ([]byte, error) {
	v := make([]byte, size)
	e := p.read(addr, &v)
	return v, e
}

func (p *Process) WriteInt8(addr uintptr, v int8) error {
	return p.write(addr, &v)
}

func (p *Process) WriteInt16(addr uintptr, v int16) error {
	return p.write(addr, &v)
}

func (p *Process) WriteInt32(addr uintptr, v int32) error {
	return p.write(addr, &v)
}

func (p *Process) WriteInt64(addr uintptr, v int64) error {
	return p.write(addr, &v)
}

func (p *Process) WriteUint8(addr uintptr, v uint8) error {
	return p.write(addr, &v)
}

func (p *Process) WriteUint16(addr uintptr, v uint16) error {
	return p.write(addr, &v)
}

func (p *Process) WriteUint32(addr uintptr, v uint32) error {
	return p.write(addr, &v)
}

func (p *Process) WriteUint64(addr uintptr, v uint64) error {
	return p.write(addr, &v)
}

func (p *Process) WriteFloat32(addr uintptr, v float32) error {
	return p.write(addr, &v)
}

func (p *Process) WriteFloat64(addr uintptr, v float64) error {
	return p.write(addr, &v)
}

func (p *Process) WriteBytes(addr uintptr, v []byte) error {
	return p.write(addr, &v)
}
