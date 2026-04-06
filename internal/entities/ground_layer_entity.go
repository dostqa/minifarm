package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"
)

type GroundLayer struct {
	PositionComponent
	SingleSpriteComponent
}

func NewGroundLayer(assetStorage *storage.AssetStorage) *Tree {
	t := &Tree{
		PositionComponent: PositionComponent{
			position: gametypes.Point{0, 0},
		},
		SingleSpriteComponent: SingleSpriteComponent{
			storage: storage.DefaultAssetStorage,
			animationInfo: animationInfo{
				id:          "groundlayer",
				frameCount:  1,
				frameWidth:  256,
				frameHeight: 144,
			},
		},
	}

	// подключение к хранилищу ресурсов не по умолчанию
	if assetStorage != nil {
		t.SingleSpriteComponent.storage = assetStorage
	}

	return t
}
