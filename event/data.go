package event

type TopicEventData interface {
	TopicID() TopicID
	EventID() EventID
	ObjectHeader() ObjectHeader
	Data() []byte
}
