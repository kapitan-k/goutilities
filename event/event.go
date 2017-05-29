package event

type Event interface {
	TopicID() TopicID
	EventID() EventID
	Time() int64
}

type IteratedEvent interface {
	FromKV(k, v []byte) (topicID TopicID, eventID []byte, kRet, vRet []byte)
}

type EventBufferIterator interface {
	NextEvent() (topicID TopicID, k, v []byte)
	Reset()
}
