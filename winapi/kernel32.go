package winapi

import (
	"syscall"
	"unsafe"
)

var (
	kernel32                   = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex            = kernel32.NewProc("CreateMutexW")
	procOpenMutex              = kernel32.NewProc("OpenMutexW")
	getSystemDefaultLocaleName = kernel32.NewProc("GetSystemDefaultLocaleName")
	getSystemDefaultLCID       = kernel32.NewProc("GetSystemDefaultLCID")
)

func GetSystemDefaultLocaleName() string {
	lcid, _, _ := getSystemDefaultLCID.Call()
	buf := make([]uint16, 128)
	_, _, _ = getSystemDefaultLocaleName.Call(uintptr(unsafe.Pointer(&buf[0])), lcid)
	return syscall.UTF16ToString(buf)
}

//CreateMutex kernel32 API CreateMutex
func CreateMutex(name string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))

	ret, _, err := procCreateMutex.Call(0, 0, paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

//OpenMutex kernel32 API OpenMutex
// 2031617 SYNCRONIZE
func OpenMutex(dwDesiredAccess uintptr, bInheritHandle uintptr, lpName string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpName)))
	ret, _, err := procOpenMutex.Call(
		dwDesiredAccess, //0x00100000, // SYNCRONIZE
		bInheritHandle,  // Not inheritable
		paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}
