package main

import (
	"minifarm/internal/commands"
	"minifarm/internal/entities"
	"minifarm/internal/input"
	"minifarm/internal/render"
	"minifarm/internal/ticker"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *entities.Player
}

func NewGame() *Game {
	player := entities.NewPlayer(nil, nil)

	return &Game{
		player: player,
	}
}

func (g *Game) Update() error {
	ticker.DefaultTicker.Update()

	err := input.DefaultInput.HandleInput(g.player)
	commands.DefaultInvoker.ExecuteCommmands()

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DefaultRender.Render(screen, g.player)
}

func (g *Game) Layout(_, _ int) (int, int) { return 600, 600 }
