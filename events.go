package main

type Event struct {
	name string
	text string
	err  error
}

func NewEvent(name, text string, err error) *Event {
	return &Event{
		name: name,
		text: text,
		err:  err,
	}
}

type EventStore struct {
	events []*Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: []*Event{},
	}
}
func (e *EventStore) AddEvent(event *Event) {
	e.events = append(e.events, event)
}

func (e *EventStore) GetEvents() []*Event {
	if e.events == nil {
		return []*Event{}
	}
	return e.events
}
