package entities

import (
	"minifarm/internal/events"
)

// Publisher - издатель.
// Любая сущность с этим компонентом может публиковать (издавать) игровые события
type Publisher interface {
	Publish(events.Event)
}
