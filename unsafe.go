package goutilities

import (
	"reflect"
	"unsafe"
)

func UintptrToByteSlice(ptr uintptr, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

func UnsafeToByteSlice(ptr unsafe.Pointer, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

func ByteSliceToUintptr(data []byte) uintptr {
	return uintptr(unsafe.Pointer(&data[0]))
}

func ByteSliceToUnsafe(data []byte) unsafe.Pointer {
	return unsafe.Pointer(&data[0])
}

func UnsafeToUint64Slice(ptr unsafe.Pointer, sz uint64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

func UnsafeToInt64Slice(ptr unsafe.Pointer, sz uint64) []int64 {
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}
