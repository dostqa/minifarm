package entities

import (
	"minifarm/internal/gametypes"
	"minifarm/internal/storage"

	"github.com/hajimehoshi/ebiten/v2"
)

// state - интерфейс, который отдаёт данные, связанные с движением
type state interface {
	IsIdle() bool
	Facing() gametypes.Vector
}

// SingleSpriteComponent - используется для статичных сущностей
type SingleSpriteComponent struct {
	storage *storage.AssetStorage

	animationInfo
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *SingleSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetSprite(s)
}

// DirectionalSpriteComponent - используется для сущностей,
// способных двигаться в пространстве
type DirectionalSpriteComponent struct {
	state
	storage *storage.AssetStorage

	animationInfo
}

// Sprite - частично реализует поведение интерфейса render.Spriter,
// отдаёт на рендер спрайт.
func (s *DirectionalSpriteComponent) Sprite() *ebiten.Image {
	return s.storage.GetSprite(s)
}

// animationInfo хранит информацию об используемой анимации.
type animationInfo struct {
	id          string
	frameCount  int
	frameWidth  int
	frameHeight int
}

// методы ниже реализует поведение
// интерфейса SpriteInfoProvider
// из пакета storage

func (c *animationInfo) ID() string {
	return c.id
}

func (c *animationInfo) FrameCount() int {
	return c.frameCount
}

func (c *animationInfo) FrameWidth() int {
	return c.frameWidth
}

func (c *animationInfo) FrameHeight() int {
	return c.frameHeight
}
