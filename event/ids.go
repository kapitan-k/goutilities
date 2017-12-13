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

type TopicEventEventKey struct {
	TopicID TopicID
	EventID EventID
}

func (self *TopicEventEventKey) AsSlice() (data []byte) {
	return UnsafeToSlice(unsafe.Pointer(self), TopicEventEventKeyByteSz)
}

func (self *TopicEventEventKey) Slice() (data []byte) {
	return CopyBuf(UnsafeToSlice(unsafe.Pointer(self), TopicEventEventKeyByteSz))
}

func (self *TopicEventEventKey) FromSlice(data []byte) {
	*self = *(*TopicEventEventKey)(unsafe.Pointer(&data[0]))
}

func TopicEventEventKeyMax(topicID TopicID) (self TopicEventEventKey) {
	self.TopicID = topicID
	self.EventID = EventIDMax
	return
}
