package event

type FnBaseEventDataToTopicEvent func(bed *BaseEventData) (te TopicEvent)
type FnDataToTopicEvent func(data []byte) (te TopicEvent)
type FnDataToTopicEvents func(data []byte) (tes []TopicEvent)

type Event interface {
	EventID() EventID
	Time() int64
	EventType() uint64
}

type TopicEvent interface {
	Event
	TopicID() TopicID
	CopyTo(other TopicEvent)
}

type FullDataTopicEvent interface {
	TopicEvent
	Data() []byte
}
