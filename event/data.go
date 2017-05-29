package event

import (
	. "github.com/kapitan-k/goutilities/data/binary"
	. "goutilities/data/binary"
	"unsafe"
)

type EventValue []byte

func EventValueCreate(minSize uint64) EventValue {
	return EventValue(make([]byte, EventHeaderByteSz+minSize))
}

func (self EventValue) EventHeader() EventHeader {
	return *(*EventHeader)(unsafe.Pointer(&self[0]))
}

func (self EventValue) Data() []byte {
	return self[EventHeaderByteSz:]
}

func (self EventValue) SetEventHeader(eh EventHeader) {
	*(*EventHeader)(unsafe.Pointer(&self[0])) = eh
}

type EventValueBuffer struct {
	ArrayLenPrefixedBuffer
}

func ByteSlicesToEventValues(datas [][]byte) []EventValue {
	return *(*[]EventValue)(unsafe.Pointer(&datas))
}

func EventValuesToByteSlices(datas []EventValue) [][]byte {
	return *(*[][]byte)(unsafe.Pointer(&datas))
}
