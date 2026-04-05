package world

import (
	"minifarm/internal/entities"
	"minifarm/internal/gametypes"

	"github.com/hajimehoshi/ebiten/v2"
)

type entity interface{}

type renderable interface {
	Sprite() *ebiten.Image
	X() float64
	Y() float64
}

type World struct {
	player      *entities.Player
	entities    []entity
	renderables []renderable
}

func (w *World) Add(e entity) {
	w.entities = append(w.entities, e)

	if r, ok := e.(renderable); ok {
		w.renderables = append(w.renderables, r)
	}
}

func (w *World) SpawnPlayer() {
	w.player = entities.NewPlayer(nil, nil)
	w.Add(w.player)
}

func (w *World) SpawnTree() {
	tree := entities.NewTree(gametypes.Point{0, 0}, nil)
	w.Add(tree)
}

func (w *World) Player() *entities.Player {
	return w.player
}

func (w *World) Renderables() []renderable {
	return w.renderables
}
