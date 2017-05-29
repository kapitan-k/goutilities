package event

const EventHeaderByteSz = 16

// First uint32 corresponds to ArrayLen
type EventHeader struct {
	ObjectHeader ObjectHeader
	EventID      EventID
}

const ObjectHeaderByteSz = 8

type ObjectHeader struct {
	Size         uint32
	InternalType uint8
	ExternalType uint8
	Format       uint8

	pad uint8
}
