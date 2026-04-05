package main

import (
	"minifarm/internal/commands"
	"minifarm/internal/input"
	"minifarm/internal/render"
	"minifarm/internal/ticker"
	"minifarm/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world *world.World
}

func NewGame() *Game {
	world := &world.World{}
	world.SpawnPlayer()
	world.SpawnTree()

	return &Game{
		world: world,
	}
}

func (g *Game) Update() error {
	ticker.DefaultTicker.Update()

	err := input.DefaultInput.HandleInput(g.world.Player())
	commands.DefaultInvoker.ExecuteCommmands()

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, r := range g.world.Renderables() {
		render.DefaultRender.Render(screen, r)
	}

}

func (g *Game) Layout(_, _ int) (int, int) { return 600, 600 }
