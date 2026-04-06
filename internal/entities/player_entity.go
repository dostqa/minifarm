package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"
	"minifarm/internal/ticker"
)

// Composition Player
type Player struct {
	Publisher

	PositionComponent
	MoveComponent
	StateComponent

	ToolbarComponent
	DirectionalSpriteComponent
}

func NewPlayer(customEventBus *events.Bus, customAssetStorage *storage.AssetStorage, customTicker *ticker.Ticker) *Player {
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
			id:      "player",
		},
	}

	// прокидываем связь между компонентами
	p.MoveComponent.Publisher = p.Publisher
	p.MoveComponent.PositionComponent = &p.PositionComponent
	p.MoveComponent.StateSetter = &p.StateComponent

	p.ToolbarComponent.Publisher = p.Publisher
	p.ToolbarComponent.PositionComponent = &p.PositionComponent

	p.StateComponent.Ticker = ticker.DefaultTicker

	p.DirectionalSpriteComponent.FacingGetter = &p.MoveComponent
	p.DirectionalSpriteComponent.StateGetter = &p.StateComponent

	// подключение к шине событий не по умолчанию
	if customEventBus != nil {
		p.Publisher = customEventBus
	}

	// подключение к хранилищу ресурсов не по умолчанию
	if customAssetStorage != nil {
		p.DirectionalSpriteComponent.storage = customAssetStorage
	}

	// подключение к тикеру не по умолчанию
	if customTicker != nil {
		p.StateComponent.Ticker = customTicker
	}

	return p
}
