package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"
)

// Composition Player
type Player struct {
	Publisher

	PositionComponent
	MoveComponent

	ToolbarComponent
	DirectionalSpriteComponent
}

func NewPlayer(bus *events.Bus, assetStorage *storage.AssetStorage) *Player {
	p := &Player{
		Publisher: &events.DefaultBus,
		MoveComponent: MoveComponent{
			step:   1,
			facing: gametypes.DownVector,
		},
		ToolbarComponent: ToolbarComponent{
			tools: [5]Tool{&Axe{}, &Axe{}, &Axe{}, &Axe{}, &Axe{}},
		},
		DirectionalSpriteComponent: DirectionalSpriteComponent{
			storage: storage.DefaultAssetStorage,
			animationInfo: animationInfo{
				id:          "player",
				frameCount:  4,
				frameWidth:  16,
				frameHeight: 16,
			},
		},
	}

	// прокидываем связь между компонентами
	p.MoveComponent.PositionComponent = &p.PositionComponent
	p.ToolbarComponent.PositionComponent = &p.PositionComponent
	p.DirectionalSpriteComponent.state = &p.MoveComponent

	p.ToolbarComponent.Publisher = p.Publisher
	p.MoveComponent.Publisher = p.Publisher

	// подключение к шине событий не по умолчанию
	if bus != nil {
		p.Publisher = bus
	}

	// подключение к хранилищу ресурсов не по умолчанию
	if assetStorage != nil {
		p.DirectionalSpriteComponent.storage = assetStorage
	}

	return p
}
