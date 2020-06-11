package types

import (
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// HeightsQueue is a simple type alias for a (buffered) channel of block heights.
type HeightsQueue chan int64

// NewHeightsQueue allows to create a new HeightsQueue containing at most size heights at a time
func NewHeightsQueue(size int) HeightsQueue {
	return make(chan int64, size)
}

// HeightsQueue is a simple type alias for a (buffered) channel of events.
type EventsQueue chan tmctypes.ResultEvent

// NewEventsQueue allows to create a new EventsQueue containing at most size events at a time
func NewEventsQueue(size int) EventsQueue {
	return make(chan tmctypes.ResultEvent, size)
}
