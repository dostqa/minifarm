package entities

import "minifarm/internal/events"

// Composition Player
type Player struct {
	Publisher

	PositionComponent
	MoveComponent

	ToolbarComponent
}

func NewPlayer(bus *events.Bus) *Player {
	p := &Player{
		Publisher: &events.DefaultBus,
		MoveComponent: MoveComponent{
			step: 3,
		},
		ToolbarComponent: ToolbarComponent{
			tools: [5]Tool{&Axe{}, &Axe{}, &Axe{}, &Axe{}, &Axe{}},
		},
	}

	// прокидываем связь между компонентами
	p.MoveComponent.PositionComponent = &p.PositionComponent
	p.ToolbarComponent.PositionComponent = &p.PositionComponent
	p.ToolbarComponent.Publisher = p.Publisher

	// подключение к шине событий не по умолчанию
	if bus != nil {
		p.Publisher = bus
	}

	return p
}
