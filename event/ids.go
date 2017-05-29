package event

import (
	. "goutilities"
	"unsafe"
)

const EventIDByteSz = 8
const TopicIDByteSz = 8

type EventID uint64
type TopicID uint64

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
		keys[i] = UintptrToByteSlice(ptr, 8)
		ptr += EventIDByteSz
	}

	return
}
