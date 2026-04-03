package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"

	"github.com/hajimehoshi/ebiten/v2"
)

// state - интерфейс, который отдаёт данные, связанные с движением
type state interface {
	isIdle() bool
	facingVector() gametypes.Vector
}

type DirectionalSpriteComponent struct {
	state   state
	storage *storage.AssetStorage

	id string
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *DirectionalSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetDirectionalSprite(s.id, s.state.isIdle(), s.state.facingVector())
}

type SingleSpriteComponent struct {
	storage *storage.AssetStorage

	id string
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *SingleSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetSingleSprite(s.id)
}
