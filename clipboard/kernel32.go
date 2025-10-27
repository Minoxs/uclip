package clipboard

import (
	"syscall"
)

var (
	kernel32     = syscall.MustLoadDLL("kernel32")
	globalAlloc  = kernel32.MustFindProc("GlobalAlloc")
	globalLock   = kernel32.MustFindProc("GlobalLock")
	globalUnlock = kernel32.MustFindProc("GlobalUnlock")
	memMove      = kernel32.MustFindProc("RtlMoveMemory")
)

type AllocFlags uint32
type Handle uintptr
type HGlobal Handle

const (
	GMEM_FIXED    AllocFlags = 0x00
	GMEM_MOVEABLE            = 0x02
	GMEM_ZEROINIT            = 0x40
)

func GlobalAlloc(flags AllocFlags, size uint) HGlobal {
	var r1, _, err = globalAlloc.Call(uintptr(flags), uintptr(size))
	if r1 == 0 {
		panic(err)
	}
	return HGlobal(r1)
}

func GlobalLock(mem HGlobal) uintptr {
	var r1, _, err = globalLock.Call(uintptr(mem))
	if r1 == 0 {
		panic(err)
	}
	return r1
}

func GlobalUnlock(mem HGlobal) {
	_, _, _ = globalUnlock.Call(uintptr(mem))
}

func StringAlloc(value string) []uint16 {
	var ret, err = syscall.UTF16FromString(value)
	if err != nil {
		panic(err)
	}
	return ret
}

func MemMove(source uintptr, destination uintptr, size uint) {
	_, _, _ = memMove.Call(destination, source, uintptr(size))
}
