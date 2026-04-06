package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateGetter interface {
	GetState() gametypes.StateName
}

// SingleSpriteComponent - используется для статичных сущностей
type SingleSpriteComponent struct {
	storage *storage.AssetStorage

	id string
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *SingleSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetSprite(s.id, gametypes.IdleStateName, gametypes.ZeroVector)
}

type FacingGetter interface {
	Facing() gametypes.Vector
}

// DirectionalSpriteComponent - используется для сущностей,
// способных двигаться в пространстве
type DirectionalSpriteComponent struct {
	StateGetter
	FacingGetter
	storage *storage.AssetStorage

	id string
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *DirectionalSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetSprite(s.id, s.GetState(), s.Facing())
}
