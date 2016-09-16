package w32

import (
	"syscall"
	"unsafe"
)

var (
	psapi = syscall.NewLazyDLL("psapi.dll")

	// Process enumeration
	pEnumProcesses = psapi.NewProc("EnumProcesses")

	// Other
	pGetProcessImageFileName = psapi.NewProc("GetProcessImageFileNameA")
)

func EnumProcesses(pProcessIds []uint32, cb uint32, pBytesReturned *uint32) bool {
	ret, _, _ := pEnumProcesses.Call(uintptr(unsafe.Pointer(&pProcessIds[0])), uintptr(cb), uintptr(unsafe.Pointer(pBytesReturned)))
	return ret != 0
}

func GetProcessImageFileName(hProcess HANDLE) (string, bool) {
	imageFileName := make([]byte, 2048)
	ret, _, _ := pGetProcessImageFileName.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&imageFileName[0])), uintptr(len(imageFileName)))
	if ret != 0 {
		return string(imageFileName[:ret]), ret != 0
	} else {
		return "", ret != 0
	}
}
