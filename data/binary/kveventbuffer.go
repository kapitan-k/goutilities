package binary

import (
	. "github.com/kapitan-k/goutilities/unsafe"
	"log"
	"unsafe"
)

const KVEventBufferHeaderByteSz = 16

type KVEventBufferHeader struct {
	FixedKeySize uint16
	Compression  uint8
	pad          uint8
	pad2         uint32

	ValueOffs uint64
	// binary representation ends here
	CntValues uint64
}

const EventBufferBaseByteSz = ArrayLenByteSz + KVEventBufferHeaderByteSz + ArrayLenByteSz + ArrayLenPrefixedBufferBaseByteSz

// KVEventBuffer is used to store Key Value data in one single byte slice.
// Binary Representation:
// ArrayLenByteSize for length of offset data
// offset data
// KVEventBufferHeader (16 bytes)
// then for keys either:
// FixedPrefixSizeBuffer
// or:
// ArrayLenPrefixedBuffer
// then for values:
// ArrayLenPrefixedBuffer
type KVEventBuffer []byte

func KVEventBufferCreate(minSize uint64) KVEventBuffer {
	return KVEventBuffer(make([]byte, EventBufferBaseByteSz+minSize))
}

func (self KVEventBuffer) KVEventBufferHeader() *KVEventBufferHeader {
	return (*KVEventBufferHeader)(unsafe.Pointer(&self[ArrayLenByteSz+self.OffsetDataSize()]))
}

func (self KVEventBuffer) OffsetData() []byte {
	return UnsafeToSlice(unsafe.Pointer(&self[ArrayLenByteSz]), uint64(GetArrayLenAt(self, 0)))
}

func (self KVEventBuffer) OffsetDataSize() uint64 {
	return uint64(GetArrayLenAt(self, 0))
}

func (self KVEventBuffer) Data() []byte {
	return self[ArrayLenByteSz+self.OffsetDataSize()+KVEventBufferHeaderByteSz+ArrayLenByteSz:]
}

func (self KVEventBuffer) Cnt() uint64 {
	return uint64(GetArrayLenAt(self, ArrayLenByteSz+self.OffsetDataSize()+KVEventBufferHeaderByteSz))
}

func (self KVEventBuffer) SetCnt(cnt uint64) {
	SetArrayLenAt(self, ArrayLenByteSz+self.OffsetDataSize()+KVEventBufferHeaderByteSz, uint32(cnt))
}

func (self *KVEventBuffer) FromHolder(h *KVEventBufferHolder, bufferBeginPad uint64) {
	*self = KVEventBufferHolderToEventBuffer(h, bufferBeginPad)
}

func KVEventBufferHeaderToFlat(data []byte, self *KVEventBufferHeader) (l uint64) {
	*(*KVEventBufferHeader)(unsafe.Pointer(&data[0])) = *self
	return KVEventBufferHeaderByteSz
}

func KVEventBufferHeaderFromFlat(self *KVEventBufferHeader, data []byte) (l uint64) {
	*self = *(*KVEventBufferHeader)(unsafe.Pointer(&data[0]))
	return KVEventBufferHeaderByteSz
}

type KVEventBufferHolder struct {
	Keb            KVEventBufferHeader
	Keys           [][]byte
	Values         [][]byte
	OffsetData     []byte
	BufferBeginPad uint64
}

func KVEventBufferHolderFromEventBuffer(self *KVEventBufferHolder, evb KVEventBuffer) {
	keb := &self.Keb
	*keb = *evb.KVEventBufferHeader()

	log.Println("KVEventBufferHeader ValueOffs", keb.ValueOffs, evb.KVEventBufferHeader())
	ptr := uintptr(unsafe.Pointer(&evb[0]))
	self.OffsetData = evb.OffsetData()
	log.Println("offset data", len(self.OffsetData), GetArrayLenAt(evb, 0))
	fpb := FixedPrefixSizeBuffer(evb[ArrayLenByteSz+len(self.OffsetData)+KVEventBufferHeaderByteSz : keb.ValueOffs])

	if keb.FixedKeySize > 0 {
		self.Keys = fpb.Datas()
	} else {
		log.Panicln("different sized keys are not supported yet")
		self.Keys, _ = ArrayLenPrefixedDataToDatasByUintptr(ptr)
	}

	vals, _ := ArrayLenPrefixedDataToDatasByUintptr(ptr + uintptr(keb.ValueOffs))
	self.Values = vals

	keb.CntValues = fpb.Cnt()

}

func KVEventBufferHolderToEventBufferDirect(self *KVEventBufferHolder) []byte {
	return KVEventBufferHolderToEventBuffer(self, self.BufferBeginPad)
}

// KVEventBufferHolderToEventBuffer creates a new byte slice from KVEventBufferHolder with the KVEventBuffer starting at bufferBeginPad.
func KVEventBufferHolderToEventBuffer(self *KVEventBufferHolder, bufferBeginPad uint64) []byte {
	evbBuf := KVEventBufferHolderToEventBufferOffsetOnly(self, uint64(len(self.OffsetData)), bufferBeginPad)
	evb := KVEventBuffer(evbBuf[bufferBeginPad:])
	copy(evb.OffsetData(), self.OffsetData)
	return evbBuf
}

// KVEventBufferHolderToEventBufferOffsetOnly creates a new byte slice with the KVEventBuffer starting at bufferBeginPad.
func KVEventBufferHolderToEventBufferOffsetOnly(self *KVEventBufferHolder, offsetDataLen uint64, bufferBeginPad uint64) []byte {
	var l, vl uint64
	fixedKeySize := self.Keb.FixedKeySize
	if fixedKeySize == 0 {
		log.Panicln("fixedKeySize == 0")
	}
	keys := self.Keys
	values := self.Values
	offsetData := self.OffsetData
	log.Println("offset data self", self.OffsetData, bufferBeginPad, fixedKeySize)

	offsetDataLen += ArrayLenByteSz

	l = uint64(len(offsetData))
	l += uint64(fixedKeySize) * uint64(len(keys))
	for _, value := range values {
		vl += uint64(len(value) + ArrayLenByteSz)
	}
	l += vl

	evbBuf := []byte(KVEventBufferCreate(bufferBeginPad + l))
	evb := KVEventBuffer(evbBuf[bufferBeginPad:])

	if len(offsetData) > 0 {
		SetArrayLenAt([]byte(evb), 0, uint32(len(offsetData)))
	}

	ebh := evb.KVEventBufferHeader()
	ebh.FixedKeySize = fixedKeySize
	ebh.ValueOffs = offsetDataLen + KVEventBufferHeaderByteSz +
		ArrayLenByteSz + uint64(fixedKeySize)*uint64(len(keys))

	log.Println("offset")

	if len(keys) == 0 {
		return evbBuf
	}

	fpb := FixedPrefixSizeBuffer(UintptrToSlice(uintptr(unsafe.Pointer(&evb[offsetDataLen+KVEventBufferHeaderByteSz])), uint64(fixedKeySize)*uint64(len(keys))+ArrayLenByteSz))
	fpb.SetDatas(keys)

	apb := ArrayLenPrefixedBuffer(UintptrToSlice(uintptr(unsafe.Pointer(&evb[ebh.ValueOffs])), vl))
	apb.SetDatas(values)

	return evbBuf
}

type KVEventBufferHolderIterator struct {
	h   *KVEventBufferHolder
	pos int
}

func KVEventBufferHolderIteratorCreate(h *KVEventBufferHolder) (self KVEventBufferHolderIterator) {
	self.h = h
	return
}

func (self *KVEventBufferHolderIterator) Next() {
	self.pos++
}

func (self *KVEventBufferHolderIterator) Valid() bool {
	return self.pos < len(self.h.Keys)
}

func (self *KVEventBufferHolderIterator) KeyValue() (k, v []byte) {
	h := self.h
	pos := self.pos
	return h.Keys[pos], h.Values[pos]
}
