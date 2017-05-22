package goutilities

import (
	"reflect"
	"unsafe"
)

func UintptrToByteSlice(ptr uintptr, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  int(sz),
		Cap:  int(sz),
		Data: uintptr(ptr),
	}))
}

func UnsafeToByteSlice(ptr unsafe.Pointer, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  int(sz),
		Cap:  int(sz),
		Data: uintptr(ptr),
	}))
}

func ByteSliceToUintptr(data []byte) uintptr {
	return uintptr(unsafe.Pointer(&data[0]))
}

func ByteSliceToUnsafe(data []byte) unsafe.Pointer {
	return unsafe.Pointer(&data[0])
}
