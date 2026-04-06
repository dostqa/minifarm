package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
)

type StateSetter interface {
	SetState(gametypes.StateName)
}

// MoveComponent - компонент движения.
// Любая сущность с этим компонентом может изменять (двигать) свою позицию на карте
type MoveComponent struct {
	Publisher
	StateSetter

	*PositionComponent

	step   float64 // длина шага за одно нажатие клавиши
	facing gametypes.Vector
}

func (m *MoveComponent) Move(direction gametypes.Vector) {
	m.SetY(m.Y() + (direction[1] * m.step))
	m.SetX(m.X() + (direction[0] * m.step))

	m.facing = direction
	m.SetState(gametypes.MoveStateName)

	m.Publish(events.NewEntityMovedEvent())
}

func (m *MoveComponent) Stop() {
	m.SetState(gametypes.IdleStateName)
}

func (m *MoveComponent) Facing() gametypes.Vector {
	return m.facing
}
