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
	world.SpawnGroundLayer()
	world.SpawnPlayer()
	world.SpawnTree(1, 1)

	return &Game{
		world: world,
	}
}

func (g *Game) Update() error {
	ticker.DefaultTicker.Update()

	err := input.DefaultInput.HandleInput(g.world.Player())
	commands.DefaultInvoker.ExecuteCommmands()

	g.world.SortRenderables()

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.Offscreen.Clear()

	for _, r := range g.world.GroundLayer() {
		render.DefaultRender.Render(render.Offscreen, r)
	}

	for _, r := range g.world.Renderables() {
		render.DefaultRender.Render(render.Offscreen, r)
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(render.ScaleValue, render.ScaleValue)

	screen.DrawImage(render.Offscreen, options)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return render.Width * render.ScaleValue, render.Height * render.ScaleValue
}
