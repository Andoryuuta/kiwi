package w32

import (
	"syscall"
	"unsafe"
)

var (
	k32 = syscall.NewLazyDLL("kernel32.dll")

	// Read / Write mem
	pReadProcessMemory  = k32.NewProc("ReadProcessMemory")
	pWriteProcessMemory = k32.NewProc("WriteProcessMemory")

	// Process enumeration
	pOpenProcess              = k32.NewProc("OpenProcess")
	pCreateToolhelp32Snapshot = k32.NewProc("CreateToolhelp32Snapshot")
	pModule32First            = k32.NewProc("Module32FirstW")
	pModule32Next             = k32.NewProc("Module32NextW")

	// Other
	pCloseHandle = k32.NewProc("CloseHandle")
)

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress, lpBuffer uintptr, nSize uintptr) (uintptr, bool) {
	var bytesRead uintptr
	ret, _, _ := pReadProcessMemory.Call(uintptr(hProcess), uintptr(lpBaseAddress), uintptr(lpBuffer), nSize, uintptr(unsafe.Pointer(&bytesRead)))
	return bytesRead, ret != 0
}

func WriteProcessMemory(hProcess HANDLE, lpBaseAddress, lpBuffer uintptr, nSize uintptr) (uintptr, bool) {
	var bytesWritten uintptr
	ret, _, _ := pWriteProcessMemory.Call(uintptr(hProcess), uintptr(lpBaseAddress), uintptr(lpBuffer), nSize, uintptr(unsafe.Pointer(&bytesWritten)))
	return bytesWritten, ret != 0
}

func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, processId uint32) (HANDLE, bool) {
	ret, _, _ := pOpenProcess.Call(uintptr(dwDesiredAccess), uintptr(*(*byte)(unsafe.Pointer(&bInheritHandle))), uintptr(processId))
	return HANDLE(ret), ret != 0
}

func CreateToolhelp32Snapshot(dwFlags uint32, th32ProcessID uint32) (HANDLE, bool) {
	ret, _, _ := pCreateToolhelp32Snapshot.Call(uintptr(dwFlags), uintptr(th32ProcessID))
	return HANDLE(ret), int(ret) != INVALID_HANDLE_VALUE
}

func Module32First(hSnapshot HANDLE, lpme *MODULEENTRY32) bool {
	ret, _, _ := pModule32First.Call(uintptr(hSnapshot), uintptr(unsafe.Pointer(&lpme)))
	return ret == 1
}

func Module32Next(hSnapshot HANDLE, lpme *MODULEENTRY32) bool {
	ret, _, _ := pModule32Next.Call(uintptr(hSnapshot), uintptr(unsafe.Pointer(&lpme)))
	return ret == 1
}

func CloseHandle(hObject HANDLE) bool {
	ret, _, _ := pCloseHandle.Call(uintptr(hObject))
	return ret != 0
}
