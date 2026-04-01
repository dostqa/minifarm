package events

import "minifarm/internal/gametypes"

type EntityMovedEvent struct{}

func (e *EntityMovedEvent) Type() gametypes.EventType {
	return gametypes.EntityMovedEventType
}

func NewEntityMovedEvent() *EntityMovedEvent {
	return &EntityMovedEvent{}
}
