package clipboard

import (
	"syscall"
	"unsafe"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	openClipboard    = user32.MustFindProc("OpenClipboard")
	closeClipboard   = user32.MustFindProc("CloseClipboard")
	setClipboardData = user32.MustFindProc("SetClipboardData")
)

func Open() error {
	var r1, _, err = openClipboard.Call(0)
	if r1 == 0 {
		return err
	}
	return nil
}

func Close() error {
	var r1, _, err = closeClipboard.Call()
	if r1 == 0 {
		return err
	}
	return nil
}

func SetClipboardData(value string) error {
	var (
		source          = StringAlloc(value)
		sourceLen       = int(unsafe.Sizeof(source[0])) * len(source)
		clipboardMemory = GlobalAlloc(GMEM_MOVEABLE|GMEM_ZEROINIT, uint(sourceLen))
		destination     = GlobalLock(clipboardMemory)
	)
	defer GlobalUnlock(clipboardMemory)
	MemMove(uintptr(unsafe.Pointer(&source[0])), destination, uint(sourceLen))

	const CF_UNICODETEXT = 13
	var r1, _, err = setClipboardData.Call(CF_UNICODETEXT, destination)
	if r1 == 0 {
		return err
	}
	return nil
}
