package kiwi

//TODO: switch from
//github.com/Andoryuuta/w32 to github.com/VividCortex/w32 if pull request is accepted.
import (
	"errors"
	"fmt"
	"github.com/Andoryuuta/w32"
	"path/filepath"
	"reflect"
	"unsafe"
)

type Process struct {
	Handle w32.HANDLE
	PID    uint32
}

const (
	PROCESS_ALL_ACCESS = w32.PROCESS_VM_READ | w32.PROCESS_VM_WRITE | w32.PROCESS_VM_OPERATION | w32.PROCESS_QUERY_INFORMATION | w32.PROCESS_CREATE_THREAD
)

//getFileNameByPID returns a file name given a PID.
func getFileNameByPID(pid uint32) string {
	var fileName string = `<Unknown File>`

	//Open process
	hnd, ok := w32.OpenProcess(w32.PROCESS_QUERY_INFORMATION, false, pid)
	if !ok {
		return fileName
	}
	defer w32.CloseHandle(hnd)

	//Get file path
	path, ok := w32.GetProcessImageFileName(hnd)
	if !ok {
		return fileName
	}

	//Split file path to get file name
	_, fileName = filepath.Split(path)
	return fileName
}

func GetProcessByFileName(fileName string) (Process, error) {
	//Read in process ids
	PIDs := make([]uint32, 1024)
	var bytesRead uint32 = 0
	ok := w32.EnumProcesses(PIDs, uint32(len(PIDs)), &bytesRead)
	if !ok {
		panic("Error Enumerating processes.")
	}

	//Loop over pids,
	//Divide bytesRead by sizeof(uint32) to get how many processes there are.
	for i := uint32(0); i < (bytesRead / 4); i++ {
		//Make sure to skip over the system process with PID 0
		if PIDs[i] == 0 {
			continue
		}

		//Check if its the process
		if getFileNameByPID(PIDs[i]) == fileName {
			hnd, ok := w32.OpenProcess(PROCESS_ALL_ACCESS, false, PIDs[i])
			if !ok {
				return Process{}, errors.New(fmt.Sprintf("Error while opening process %d", PIDs[i]))
			}
			return Process{hnd, PIDs[i]}, nil
		}
	}

	//Couldn't find process, return an error
	return Process{}, errors.New("Couldn't find process with name " + fileName)
}

func (p *Process) ReadInt8(addr uintptr) (v int8, e error)       { e = p.read(addr, &v); return v, e }
func (p *Process) ReadInt16(addr uintptr) (v int16, e error)     { e = p.read(addr, &v); return v, e }
func (p *Process) ReadInt32(addr uintptr) (v int32, e error)     { e = p.read(addr, &v); return v, e }
func (p *Process) ReadInt64(addr uintptr) (v int64, e error)     { e = p.read(addr, &v); return v, e }
func (p *Process) ReadUint8(addr uintptr) (v uint8, e error)     { e = p.read(addr, &v); return v, e }
func (p *Process) ReadUint16(addr uintptr) (v uint16, e error)   { e = p.read(addr, &v); return v, e }
func (p *Process) ReadUint32(addr uintptr) (v uint32, e error)   { e = p.read(addr, &v); return v, e }
func (p *Process) ReadUint64(addr uintptr) (v uint64, e error)   { e = p.read(addr, &v); return v, e }
func (p *Process) ReadFloat32(addr uintptr) (v float32, e error) { e = p.read(addr, &v); return v, e }
func (p *Process) ReadFloat64(addr uintptr) (v float64, e error) { e = p.read(addr, &v); return v, e }
func (p *Process) ReadUintptr(addr uintptr) (v uintptr, e error) { e = p.read(addr, &v); return v, e }

func (p *Process) read(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	i := reflect.Indirect(v)
	size := i.Type().Size()
	bytesRead, ok := w32.ReadProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(i.UnsafeAddr()),
		size,
	)
	if !ok || bytesRead != size {
		return errors.New("Error on reading process memory.")
	}
	return nil
}

func (p *Process) ReadUint32Ptr(addr uintptr, offsets ...uintptr) (uint32, error) {
	curPtr, err := p.ReadUint32(addr)
	if err != nil{
        return 0, errors.New("Error while trying to read from ptr base.")
    }
	
	for _, offset := range offsets{
		fmt.Printf("Adding offset % X\n", offset)
		curPtr, err = p.ReadUint32(uintptr(curPtr)+offset)
		if err != nil{
			return 0, errors.New("Error while trying to read from offset.")
		}
	}
	
	return curPtr, nil
}

func (p *Process) WriteInt8(addr uintptr, v int8) (e error)       { return p.write(addr, &v) }
func (p *Process) WriteInt16(addr uintptr, v int16) (e error)     { return p.write(addr, &v) }
func (p *Process) WriteInt32(addr uintptr, v int32) (e error)     { return p.write(addr, &v) }
func (p *Process) WriteInt64(addr uintptr, v int64) (e error)     { return p.write(addr, &v) }
func (p *Process) WriteUint8(addr uintptr, v uint8) (e error)     { return p.write(addr, &v) }
func (p *Process) WriteUint16(addr uintptr, v uint16) (e error)   { return p.write(addr, &v) }
func (p *Process) WriteUint32(addr uintptr, v uint32) (e error)   { return p.write(addr, &v) }
func (p *Process) WriteUint64(addr uintptr, v uint64) (e error)   { return p.write(addr, &v) }
func (p *Process) WriteFloat32(addr uintptr, v float32) (e error) { return p.write(addr, &v) }
func (p *Process) WriteFloat64(addr uintptr, v float64) (e error) { return p.write(addr, &v) }
func (p *Process) WriteUintptr(addr uintptr, v uintptr) (e error) { return p.write(addr, &v) }

func (p *Process) write(addr uintptr, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	i := reflect.Indirect(v)
	size := i.Type().Size()
	bytesWritten, ok := w32.WriteProcessMemory(
		p.Handle,
		unsafe.Pointer(addr),
		unsafe.Pointer(i.UnsafeAddr()),
		size,
	)
	if !ok || bytesWritten != size {
		return errors.New("Error on writing process memory.")
	}
	return nil
}
