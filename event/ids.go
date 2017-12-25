package event

import (
	. "github.com/kapitan-k/goutilities/data"
	. "github.com/kapitan-k/goutilities/unsafe"
	"unsafe"
)

type EventIDKey [EventIDByteSz]byte

func (self *EventIDKey) EventID() EventID {
	return *(*EventID)(unsafe.Pointer(&self[0]))
}

func (self *EventIDKey) SetEventID(eid EventID) {
	*(*EventID)(unsafe.Pointer(&self[0])) = eid
}

func EventIDKeyFromEventID(eid EventID) (eik EventIDKey) {
	eik.SetEventID(eid)
	return
}

func EventIDFromBuf(data []byte) EventID {
	return *(*EventID)(unsafe.Pointer(&data[0]))
}

func EventIDsToFixedKeys(eventIDs []EventID) (keys [][]byte) {
	keys = make([][]byte, len(eventIDs))
	ptr := uintptr(unsafe.Pointer(&eventIDs[0]))

	for i := range eventIDs {
		keys[i] = UintptrToSlice(ptr, 8)
		ptr += EventIDByteSz
	}

	return
}

const TopicEventEventKeyByteSz = 16

// TopicEventEventKey defines a 16 byte unique event identification.
// first 8 byte is TopicID, second 8 byte is EventID
type TopicEventEventKey struct {
	TopicID TopicID
	EventID EventID
}

// AsSlice returns the TopicEventEventKey as a byte slice
// but unsafe!
func (self *TopicEventEventKey) AsSlice() (data []byte) {
	return UnsafeToSlice(unsafe.Pointer(self), TopicEventEventKeyByteSz)
}

// Slice returns the TopicEventEventKey as a freshly allocated byte slice.
func (self *TopicEventEventKey) Slice() (data []byte) {
	return CopyBuf(UnsafeToSlice(unsafe.Pointer(self), TopicEventEventKeyByteSz))
}

// FromSlice copies the data from data to TopicEventEventKey.
func (self *TopicEventEventKey) FromSlice(data []byte) {
	*self = *(*TopicEventEventKey)(unsafe.Pointer(&data[0]))
}

// Size implements the marshaler interface.
func (self *TopicEventEventKey) Size() int {
	return TopicEventEventKeyByteSz
}

// Marshal implements the marshaler interface.
func (self *TopicEventEventKey) Marshal() (data []byte, err error) {
	return self.Slice(), nil
}

// MarshalTo implements the marshaler interface.
func (self *TopicEventEventKey) MarshalTo(data []byte) (err error) {
	copy(data, self.AsSlice())
	return
}

// Unmarshal implements the unmarshaler interface.
func (self *TopicEventEventKey) Unmarshal(data []byte) (err error) {
	self.FromSlice(data)
	return
}

// TopicEventEventKeyMax return a TopicEventEventKey
// for topicID with a maximum EventID.
func TopicEventEventKeyMax(topicID TopicID) (self TopicEventEventKey) {
	self.TopicID = topicID
	self.EventID = EventIDMax
	return
}
