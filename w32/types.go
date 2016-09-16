package w32

type (
	HANDLE  uintptr
	HMODULE HANDLE
)

type MODULEENTRY32 struct {
	DwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   *uint8
	ModBaseSize   uint32
	HMODULE       HMODULE
	SzModule      [MAX_MODULE_NAME32 + 1]uint16
	SzExePath     [MAX_PATH]uint16
}
