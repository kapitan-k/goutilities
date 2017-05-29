package binary

import (
	. "github.com/kapitan-k/goutilities"
	"unsafe"
)

// this buffer has 4 bytes (uint32t) at the first bytes which define the count
// of its elements.
// afterwards elements must follow in |value|value
// each with the same len
type FixedPrefixSizeBuffer []byte

func FixedPrefixSizeBufferCreate(elemSz, cnt uint64) FixedPrefixSizeBuffer {
	fpb := FixedPrefixSizeBuffer(make([]byte, elemSz*cnt+ArrayLenByteSz))
	fpb.SetCnt(cnt)
	return fpb
}

func (self FixedPrefixSizeBuffer) Cnt() uint64 {
	return uint64(GetArrayLenAt(self, 0))
}

func (self FixedPrefixSizeBuffer) SetCnt(cnt uint64) {
	SetArrayLenAt(self, 0, uint32(cnt))
}

func (self FixedPrefixSizeBuffer) ElemSize() uint64 {
	return uint64(uint64(len(self)-ArrayLenByteSz) / self.Cnt())
}

func (self FixedPrefixSizeBuffer) Datas() (datas [][]byte) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	elemSz := self.ElemSize()
	datas, _ = FixedSizeDataToDatasByUintptr(ptr, elemSz)
	return
}

func (self FixedPrefixSizeBuffer) SetDatas(datas [][]byte) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	self.SetCnt(uint64(len(datas)))
	ptr += ArrayLenByteSz
	elemSz := self.ElemSize()
	for _, data := range datas {
		dst := UintptrToByteSlice(ptr, uint64(elemSz))
		copy(dst, data)
		ptr += uintptr(elemSz)
	}
}

func (self FixedPrefixSizeBuffer) Iterate(fn func(data []byte) (more bool)) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	cnt := self.Cnt()
	elemSz := self.ElemSize()
	ptr += ArrayLenByteSz
	for i := uint64(0); i < cnt; i++ {
		more := fn(UintptrToByteSlice(ptr, elemSz))
		if !more {
			return
		}
		ptr += uintptr(elemSz)
	}
}

// no bounds check
func (self FixedPrefixSizeBuffer) At(off uint64) (data []byte) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	elemSz := self.ElemSize()

	return UintptrToByteSlice(ptr+ArrayLenByteSz+uintptr(off*elemSz), elemSz)
}

// no bounds check
func (self FixedPrefixSizeBuffer) SetAt(data []byte, off uint64) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	elemSz := self.ElemSize()
	dst := UintptrToByteSlice(ptr+ArrayLenByteSz+uintptr(off*elemSz), elemSz)
	copy(dst, data)
}

func (self FixedPrefixSizeBuffer) Append(data []byte) FixedPrefixSizeBuffer {
	self = append(self, data...)
	self.SetCnt(self.Cnt() + 1)
	return self
}

const ArrayLenPrefixedBufferBaseByteSz = 4

// this buffer has 4 bytes (uint32) at the first bytes which define the count
// of its elements.
// afterwards elements must follow in valueLen|value|valueLen|value
// valueLen is represented as uint32
type ArrayLenPrefixedBuffer []byte

func ArrayLenPrefixedBufferCreate(minSize uint64) ArrayLenPrefixedBuffer {
	return ArrayLenPrefixedBuffer(make([]byte, minSize+ArrayLenByteSz))
}

func (self ArrayLenPrefixedBuffer) Cnt() uint64 {
	return uint64(GetArrayLenAt(self, 0))
}

func (self ArrayLenPrefixedBuffer) SetCnt(cnt uint64) {
	SetArrayLenAt(self, 0, uint32(cnt))
}

func (self ArrayLenPrefixedBuffer) Datas() (datas [][]byte) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	datas, _ = ArrayLenPrefixedDataToDatasByUintptr(ptr)
	return
}

// datas must be prefixed with valueLen
func (self ArrayLenPrefixedBuffer) SetDatas(datas [][]byte) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	ArrayLenPrefixedDataFromDatasByUintptr(ptr, datas)
}

func (self ArrayLenPrefixedBuffer) Iterate(fn func(data []byte) (more bool)) {
	ptr := uintptr(unsafe.Pointer(&self[0]))
	cnt := self.Cnt()
	ptr += ArrayLenByteSz
	for i := uint64(0); i < cnt; i++ {
		elemSz := uint64(GetArrayLenAtUintptr(ptr, 0))
		ptr += ArrayLenByteSz
		more := fn(UintptrToByteSlice(ptr, elemSz))
		if !more {
			return
		}
		ptr += uintptr(elemSz)
	}
}

func (self ArrayLenPrefixedBuffer) AppendArrayLenPrefixedData(data []byte) ArrayLenPrefixedBuffer {
	self = append(self, data...)
	self.SetCnt(self.Cnt() + 1)
	return self
}

// data must not be prefixed with valueLen
// very inefficient
func (self ArrayLenPrefixedBuffer) Append(data []byte) ArrayLenPrefixedBuffer {
	lb := [ArrayLenByteSz]byte{}
	SetArrayLenAt(lb[:], 0, uint32(len(data)))
	self = append(self, lb[:]...)
	self = append(self, data...)
	self.SetCnt(self.Cnt() + 1)
	return self
}

// datas must not be prefixed with valueLen
func (self ArrayLenPrefixedBuffer) AppendMulti(datas [][]byte) ArrayLenPrefixedBuffer {
	var l int
	for _, data := range datas {
		l += len(data)
	}
	lb := make([]byte, l+len(datas)*ArrayLenByteSz)
	ptr := uintptr(unsafe.Pointer(&lb[0]))
	for _, data := range datas {
		l := uint64(len(data))
		SetArrayLenAtUintptr(ptr, 0, uint32(l))
		ptr += ArrayLenByteSz
		dst := UintptrToByteSlice(ptr, l)
		copy(dst, data)
		ptr += uintptr(l)
	}

	self.SetCnt(self.Cnt() + uint64(len(datas)))

	self = append(self, lb...)

	return self
}

func ArrayLenPrefixedDataToDatasByUintptr(ptr uintptr) (datas [][]byte, ptrEnd uintptr) {
	cnt := uint64(GetArrayLenAtUintptr(ptr, 0))
	ptr += ArrayLenByteSz
	datas = make([][]byte, cnt)
	for i := uint64(0); i < cnt; i++ {
		elemSz := uint64(GetArrayLenAtUintptr(ptr, 0))
		ptr += ArrayLenByteSz
		datas[i] = UintptrToByteSlice(ptr, elemSz)
		ptr += uintptr(elemSz)
	}

	ptrEnd = ptr

	return
}

func ArrayLenPrefixedDataFromDatasByUintptr(ptr uintptr, datas [][]byte) (ptrEnd uintptr) {
	SetArrayLenAtUintptr(ptr, 0, uint32(len(datas)))
	ptr += ArrayLenByteSz
	for _, data := range datas {
		l := uint64(len(data))
		SetArrayLenAtUintptr(ptr, 0, uint32(l))
		ptr += ArrayLenByteSz
		dst := UintptrToByteSlice(ptr, l)
		copy(dst, data)
		ptr += uintptr(l)
	}
	ptrEnd = ptr

	return
}

func FixedSizeDataToDatasByUintptr(ptr uintptr, sz uint64) (datas [][]byte, ptrEnd uintptr) {
	cnt := GetArrayLenAtUintptr(ptr, 0)
	ptr += ArrayLenByteSz
	datas = make([][]byte, cnt)
	for i := uint32(0); i < cnt; i++ {
		datas[i] = UintptrToByteSlice(ptr, sz)
		ptr += uintptr(sz)
	}

	ptrEnd = ptr

	return
}
