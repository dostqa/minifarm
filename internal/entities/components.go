package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
)

// Publisher - издатель.
// Любая сущность с этим компонентом может публиковать (издавать) игровые события
type Publisher interface {
	Publish(events.Event)
}

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

// MoveComponent - компонент движения.
// Любая сущность с этим компонентом может изменять (двигать) свою позицию на карте
type MoveComponent struct {
	*PositionComponent
	step float64 // длина шага за одно нажатие клавиши
}

func (m *MoveComponent) Move(direction gametypes.Vector) {
	m.SetY(m.Y() + (direction[1] * m.step))
	m.SetX(m.X() + (direction[0] * m.step))
}

// Интерфейс описывает, что должен уметь инструмент
type Tool interface {
	Type() gametypes.ToolType
}

// ToolbarComponent - компонент панели инструментов.
// Любая сущность с этим компонентом может использовать инструменты
type ToolbarComponent struct {
	Publisher
	*PositionComponent
	tools  [5]Tool
	active int
}

func (bar *ToolbarComponent) UseActiveTool() {
	bar.Publish(events.NewToolUsedEvent(bar.Position(), bar.tools[bar.active].Type()))
}

func (bar *ToolbarComponent) ChangeActiveTool(i int) {
	// Логика, при которой происходит смена активного инструмента
	if (i >= 0) && (i <= 4) {
		if bar.active == i {
			return
		} else {
			// Непосредственно меняем активный слот
			bar.active = i

			// Создаём событие о смене активного слота
			bar.Publish(events.NewActiveToolChanged(bar.tools[bar.active].Type()))
		}
	} else {
		return
	}
}
