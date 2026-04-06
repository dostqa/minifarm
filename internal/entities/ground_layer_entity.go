package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"
)

type GroundLayer struct {
	PositionComponent
	SingleSpriteComponent
}

func NewGroundLayer(customAssetStorage *storage.AssetStorage) *GroundLayer {
	g := &GroundLayer{
		PositionComponent: PositionComponent{
			position: gametypes.Point{0, 0},
		},
		SingleSpriteComponent: SingleSpriteComponent{
			storage: storage.DefaultAssetStorage,
			id:      "groundlayer",
		},
	}

	// подключение к хранилищу ресурсов не по умолчанию
	if customAssetStorage != nil {
		g.SingleSpriteComponent.storage = customAssetStorage
	}

	return g
}
