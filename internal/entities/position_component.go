package entities

import (
	"minifarm/internal/gametypes"
)

// PositionComponent - компонент позиции.
// Любая сущность с этим компонентом может иметь позицию на карте
type PositionComponent struct {
	position gametypes.Point
}

func (p *PositionComponent) Position() gametypes.Point {
	return p.position
}

func (p *PositionComponent) X() float64 {
	return p.position[0]
}

func (p *PositionComponent) Y() float64 {
	return p.position[1]
}

func (p *PositionComponent) SetX(x float64) {
	p.position[0] = x
}

func (p *PositionComponent) SetY(y float64) {
	p.position[1] = y
}
