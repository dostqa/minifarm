package events

import "minifarm/internal/gametypes"

var (
	DefaultBus Bus
)

// Подписчик -
// Нечто, которое изьвляет желание получать
// уведомления о происходящих событиях
type Subscriber interface {
	Handle(Event)
}

// Event - описывает поведение событий
type Event interface {
	Type() gametypes.EventType
}

// Bus - глобальная шина событий
type Bus struct {
	subscribers []Subscriber
}

func (bus *Bus) Publish(event Event) {
	for _, sub := range bus.subscribers {
		sub.Handle(event)
	}
}

func (bus *Bus) Subscribe(subs ...Subscriber) {
	bus.subscribers = append(bus.subscribers, subs...)
}

func NewBus() *Bus {
	return &Bus{}
}
