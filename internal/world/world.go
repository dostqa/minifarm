package world

import (
	"minifarm/internal/entities"
	"minifarm/internal/gametypes"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type entity interface{}

type renderable interface {
	Sprite() *ebiten.Image
	X() float64
	Y() float64
}

type World struct {
	player   *entities.Player
	entities []entity

	groundLayer []renderable
	renderables []renderable
}

func (w *World) Add(e entity) {
	w.entities = append(w.entities, e)

	if r, ok := e.(renderable); ok {
		w.renderables = append(w.renderables, r)
	}
}

func (w *World) SpawnGroundLayer() {
	w.groundLayer = append(w.groundLayer, entities.NewGroundLayer(nil))
}

func (w *World) SpawnPlayer() {
	w.player = entities.NewPlayer(nil, nil)
	w.Add(w.player)
}

// SpawnTree создаёт новое дерево, на переданной позиции,
// где x и y - координаты тайла
func (w *World) SpawnTree(x, y int) {
	tree := entities.NewTree(gametypes.Point{float64(x) * 16, float64(y) * 16}, nil)
	w.Add(tree)
}

func (w *World) Player() *entities.Player {
	return w.player
}

func (w *World) SortRenderables() {
	sort.Slice(w.renderables, func(i, j int) bool { return w.renderables[i].Y() < w.renderables[j].Y() })
}

func (w *World) Renderables() []renderable {
	return w.renderables
}

func (w *World) GroundLayer() []renderable {
	return w.groundLayer
}
