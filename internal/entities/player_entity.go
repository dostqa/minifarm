package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/storage"
)

// Composition Player
type Player struct {
	Publisher

	PositionComponent
	MoveComponent

	ToolbarComponent
	SpriteComponent
}

func NewPlayer(bus *events.Bus, assetStorage *storage.AssetStorage) *Player {
	p := &Player{
		Publisher: &events.DefaultBus,
		MoveComponent: MoveComponent{
			step: 3,
		},
		ToolbarComponent: ToolbarComponent{
			tools: [5]Tool{&Axe{}, &Axe{}, &Axe{}, &Axe{}, &Axe{}},
		},
		SpriteComponent: SpriteComponent{
			storage:  storage.DefaultAssetStorage,
			spriteID: storage.PlayerSprite,
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

	// подключение к хранилищу ресурсов не по умолчанию
	if assetStorage != nil {
		p.SpriteComponent.storage = assetStorage
	}

	return p
}
