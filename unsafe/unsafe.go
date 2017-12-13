package unsafe

import (
	"reflect"
	"unsafe"
)

// UintptrToSlice returns a slice with len and cap of sz.
func UintptrToSlice(ptr uintptr, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

// UnsafeToSlice returns a slice with len and cap of sz.
func UnsafeToSlice(ptr unsafe.Pointer, sz uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

// SliceToUintptr returns the pointer to which data data points.
func SliceToUintptr(data []byte) uintptr {
	return uintptr(unsafe.Pointer(&data[0]))
}

// SliceToUnsafe returns the pointer to which data data points.
func SliceToUnsafe(data []byte) unsafe.Pointer {
	return unsafe.Pointer(&data[0])
}

//UnsafeToUint64Slice returns a uint64 slice with len and cap of sz.
func UnsafeToUint64Slice(ptr unsafe.Pointer, sz uint64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

//UnsafeToInt64Slice returns a int64 slice with len and cap of sz.
func UnsafeToInt64Slice(ptr unsafe.Pointer, sz uint64) []int64 {
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(sz),
		Cap:  int(sz),
	}))
}

// ByteSliceFromUint64 returns a byte slice with the length of 8
// and val copied to it without considering byte order.
func ByteSliceFromUint64(val uint64) []byte {
	data := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(&data[0])) = val
	return data
}

// Int64SliceToUint64Slice converts a int64 slice to a uint64 slice unsafe.
func Int64SliceToUint64Slice(vals []int64) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals),
		Cap:  cap(vals),
	}))
}

// Uint64SliceToInt64Slice converts a uint64 slice to a int64 slice unsafe.
func Uint64SliceToInt64Slice(vals []uint64) []int64 {
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals),
		Cap:  cap(vals),
	}))
}

// Uint64SliceToByteSlice converts a uint64 slice to a byte slice unsafe.
// Length of vals is multiplied by 8.
func Uint64SliceToByteSlice(vals []uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals) * 8,
		Cap:  cap(vals) * 8,
	}))
}

// ByteSliceToUint64Slice converts a byte slice to a uint64 slice unsafe.
// Length of vals is divided by 8.
func ByteSliceToUint64Slice(vals []byte) []uint64 {
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals) / 8,
		Cap:  cap(vals) / 8,
	}))
}

// Int64SliceToByteSlice converts a int64 slice to a byte slice unsafe.
// Length of vals is multiplied by 8.
func Int64SliceToByteSlice(vals []int64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals) * 8,
		Cap:  cap(vals) * 8,
	}))
}

// ByteSliceToInt64Slice converts a byte slice to a int64 slice unsafe.
// Length of vals is divided by 8.
func ByteSliceToInt64Slice(vals []byte) []int64 {
	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&vals[0])),
		Len:  len(vals) / 8,
		Cap:  cap(vals) / 8,
	}))
}
