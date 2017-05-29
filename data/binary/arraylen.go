package binary

import (
	"unsafe"
)

const ArrayLenByteSz = 4

func SetArrayLenAt(buf []byte, pos uint64, l uint32) {
	SetArrayLenAtUintptr(uintptr(unsafe.Pointer(&buf[0])), pos, l)
}

func SetArrayLenAtUintptr(ptr uintptr, pos uint64, l uint32) {
	*(*uint32)(unsafe.Pointer(ptr + uintptr(pos))) = l
}

func GetArrayLenAt(buf []byte, pos uint64) (l uint32) {
	return GetArrayLenAtUintptr(uintptr(unsafe.Pointer(&buf[0])), pos)
}

func GetArrayLenAtUintptr(ptr uintptr, pos uint64) (l uint32) {
	return *(*uint32)(unsafe.Pointer(ptr + uintptr(pos)))
}
