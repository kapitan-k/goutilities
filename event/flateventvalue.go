package event

import (
	//. "github.com/kapitan-k/goutilities/data/binary"
	"unsafe"
)

const FlatEventValueBaseByteSz = ObjectHeaderByteSz

// FlatEventValue holds a single event prefixed by its ObjectHeader.
// It can be used to send and receive a single event
type FlatEventValue []byte

// FlatEventValueCreate returns a new FlatEventValue with the byte capacity of ObjectHeaderByteSz + cap.
func FlatEventValueCreate(cap uint64) FlatEventValue {
	return make(FlatEventValue, ObjectHeaderByteSz+cap)
}

func FlatEventValueFromSlice(data []byte) FlatEventValue {
	fev := FlatEventValueCreate(uint64(len(data)))
	copy(fev.Data(), data)
	return fev
}

// FlatEventValueFromSliceWithPrefix can be used to create a larger object header
// dataPrefixSize must be calculated without ObjectHeaderByteSz
// ObjectHeader.ExtensionByteSize must reflect dataPrefixSize
func FlatEventValueFromSliceWithPrefix(data []byte, dataPrefixSize uint64) FlatEventValue {
	fev := FlatEventValueCreate(uint64(len(data)) + dataPrefixSize)
	copy(fev.Data()[dataPrefixSize:], data)
	return fev
}

func (self FlatEventValue) ObjectHeader() ObjectHeader {
	return *(*ObjectHeader)(unsafe.Pointer(&self[0]))
}

// Data returns the data without ObjectHeader.
func (self FlatEventValue) Data() []byte {
	oh := self.ObjectHeader()
	var offs int
	if oh.ExtensionByteSize == 0 {
		offs = ObjectHeaderByteSz
	} else {
		offs = int(oh.ExtensionByteSize)
	}
	return self[offs:]
}

func (self FlatEventValue) SetObjectHeader(eh ObjectHeader) {
	*(*ObjectHeader)(unsafe.Pointer(&self[0])) = eh
}

// CopyData just copies data to the Data() of this FlatEventValue, thats all
func (self FlatEventValue) CopyData(src []byte) {
	copy(self.Data(), src)
}
