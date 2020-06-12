package types

// HeightsQueue is a simple type alias for a (buffered) channel of events.
type EventsQueue chan interface{}

// NewEventsQueue allows to create a new EventsQueue containing at most size events at a time
func NewEventsQueue(size int) EventsQueue {
	return make(chan interface{}, size)
}
