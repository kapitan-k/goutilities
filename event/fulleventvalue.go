package event

import (
	"unsafe"
)

// FlatEventValue holds a single event prefixed by its FullEventHeader.
// It can be used to send and receive a single event
type FullEventValue []byte

func FullEventValueCreate(cap uint64) FullEventValue {
	return FullEventValue(make([]byte, FullEventHeaderByteSz+cap))
}

func FullEventValueFromSlice(data []byte) FullEventValue {
	fev := FullEventValueCreate(uint64(len(data)))
	copy(fev.Data(), data)
	return fev
}

func FullEventValueFromBaseEventData(bed *BaseEventData) FullEventValue {
	fev := FullEventValueCreate(uint64(len(bed.Data())))
	copy(fev.Data(), bed.Data())
	fev.SetFullEventHeader(bed.FullEventHeader())
	return fev
}

func (self FullEventValue) FullEventHeader() FullEventHeader {
	return *(*FullEventHeader)(unsafe.Pointer(&self[0]))
}

func (self FullEventValue) Data() []byte {
	return self[FullEventHeaderByteSz:]
}

func (self FullEventValue) SetFullEventHeader(eh FullEventHeader) {
	*(*FullEventHeader)(unsafe.Pointer(&self[0])) = eh
}

func FullEventValueByFlatEventValue(fev FlatEventValue) FullEventValue {
	return FullEventValue(fev)
}

/*

type FullEventValueBuffer struct {
	ArrayLenPrefixedBuffer
}

func ByteSlicesToFullEventValues(datas [][]byte) []FullEventValue {
	return *(*[]FullEventValue)(unsafe.Pointer(&datas))
}

func FullEventValuesToByteSlices(datas []FullEventValue) [][]byte {
	return *(*[][]byte)(unsafe.Pointer(&datas))
}

*/
