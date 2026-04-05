package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"
)

type Tree struct {
	PositionComponent
	SingleSpriteComponent
}

func NewTree(position gametypes.Point, assetStorage *storage.AssetStorage) *Tree {
	t := &Tree{
		PositionComponent: PositionComponent{
			position: position,
		},
		SingleSpriteComponent: SingleSpriteComponent{
			storage: storage.DefaultAssetStorage,
			id:      "tree",
		},
	}

	// подключение к хранилищу ресурсов не по умолчанию
	if assetStorage != nil {
		t.SingleSpriteComponent.storage = assetStorage
	}

	return t
}
