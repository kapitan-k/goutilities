package event

import (
	. "github.com/kapitan-k/goutilities/data"
	. "github.com/kapitan-k/goutilities/unsafe"
	"unsafe"
)

const TopicEventHeaderByteSz = TopicIDByteSz + EventIDByteSz

// TopicEventHeader holds the TopicID and the EventID of an event.
type TopicEventHeader struct {
	TopicID TopicID
	EventID EventID
}

func (self *TopicEventHeader) AsSlice() []byte {
	return UnsafeToSlice(unsafe.Pointer(self), TopicEventHeaderByteSz)
}

func (self *TopicEventHeader) Slice() []byte {
	return CopyBuf(UnsafeToSlice(unsafe.Pointer(self), TopicEventHeaderByteSz))
}

func TopicEventHeaderFromBuf(data []byte) TopicEventHeader {
	return *(*TopicEventHeader)(unsafe.Pointer(&data[0]))
}

const FullEventHeaderByteSz = ObjectHeaderByteSz + TopicEventHeaderByteSz

// FullEventHeader.
// The order of fields matters here, for example to change from Flat- to FullEventValue.
type FullEventHeader struct {
	ObjectHeader     ObjectHeader
	TopicEventHeader TopicEventHeader
}

const ObjectHeaderByteSz = 8

// ObjectHeader holds description of the type of an object or event.
type ObjectHeader struct {
	ObjectDataType     uint16
	UnderlyingDataType uint16
	Format             uint8
	Format2            uint8
	// ExtensionByteSize defines the bytes after this header
	// which belong to an implementation specific header
	ExtensionByteSize uint16
}

func ObjectHeaderFromUint64(val uint64) ObjectHeader {
	return *(*ObjectHeader)(unsafe.Pointer(&val))
}

func FullEventHeaderToKey(teh FullEventHeader) FullEventHeaderKey {
	k := FullEventHeaderKey{}
	k.SetObjectHeader(teh.ObjectHeader)
	k.SetTopicEventHeader(teh.TopicEventHeader)

	return k
}

type FullEventHeaderKey [FullEventHeaderByteSz]byte

func (self *FullEventHeaderKey) ObjectHeader() ObjectHeader {
	return *(*ObjectHeader)(unsafe.Pointer(&self[0]))
}

func (self *FullEventHeaderKey) SetObjectHeader(oh ObjectHeader) {
	*(*ObjectHeader)(unsafe.Pointer(&self[0])) = oh
}

func (self *FullEventHeaderKey) TopicEventHeader() TopicEventHeader {
	return *(*TopicEventHeader)(unsafe.Pointer(&self[0]))
}

func (self *FullEventHeaderKey) SetTopicEventHeader(tid TopicEventHeader) {
	*(*TopicEventHeader)(unsafe.Pointer(&self[ObjectHeaderByteSz])) = tid
}
