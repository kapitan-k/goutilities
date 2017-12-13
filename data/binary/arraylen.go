package binary

import (
	"unsafe"
)

// ArrayLenByteSz defines the byte size in which the length of
// an array is stored.
const ArrayLenByteSz = 4

// SetArrayLenAt sets l at the byte position of pos.
func SetArrayLenAt(buf []byte, pos uint64, l uint32) {
	SetArrayLenAtUintptr(uintptr(unsafe.Pointer(&buf[0])), pos, l)
}

// SetArrayLenAtUintptr sets l at the byte position of pos.
func SetArrayLenAtUintptr(ptr uintptr, pos uint64, l uint32) {
	*(*uint32)(unsafe.Pointer(ptr + uintptr(pos))) = l
}

// GetArrayLenAt gets l at the byte position of pos.
func GetArrayLenAt(buf []byte, pos uint64) (l uint32) {
	return GetArrayLenAtUintptr(uintptr(unsafe.Pointer(&buf[0])), pos)
}

// GetArrayLenAtUintptr gets l at the byte position of pos.
func GetArrayLenAtUintptr(ptr uintptr, pos uint64) (l uint32) {
	return *(*uint32)(unsafe.Pointer(ptr + uintptr(pos)))
}
