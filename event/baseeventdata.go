package event

// BaseEventData holds a FullEventHeader and the events data separately.
type BaseEventData struct {
	eh   FullEventHeader
	data []byte
}

func BaseEventDataCreate(topicID TopicID, eventID EventID, objectHeader ObjectHeader, data []byte) (self BaseEventData) {
	self.eh.TopicEventHeader.TopicID = topicID
	self.eh.TopicEventHeader.EventID = eventID
	self.eh.ObjectHeader = objectHeader
	self.data = data
	return
}

func (self *BaseEventData) FullEventHeader() FullEventHeader {
	return self.eh
}

func (self *BaseEventData) TopicID() (topicID TopicID) {
	return self.eh.TopicEventHeader.TopicID
}

func (self *BaseEventData) EventID() (eventID EventID) {
	return self.eh.TopicEventHeader.EventID
}

func (self *BaseEventData) ObjectHeader() ObjectHeader {
	return self.eh.ObjectHeader
}

func (self *BaseEventData) Data() []byte {
	return self.data
}

func (self *BaseEventData) SetTopicID(topicID TopicID) {
	self.eh.TopicEventHeader.TopicID = topicID
}

func (self *BaseEventData) SetEventID(eventID EventID) {
	self.eh.TopicEventHeader.EventID = eventID
}

func (self *BaseEventData) SetObjectHeader(objectHeader ObjectHeader) {
	self.eh.ObjectHeader = objectHeader
}
