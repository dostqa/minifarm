package main

import (
	"fmt"
	"minifarm/internal/commands"
	"minifarm/internal/entities"
	"minifarm/internal/gametypes"
	"minifarm/internal/input"
	"minifarm/internal/render"
	"minifarm/internal/ticker"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player  *entities.Player
	changes gametypes.Point
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

	if !((g.player.X() == g.changes[0]) && (g.player.Y() == g.changes[1])) {
		fmt.Println(g.player)
	}
	g.changes = gametypes.Point{g.player.X(), g.player.Y()}

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	render.DefaultRender.Render(screen, g.player)
}

func (g *Game) Layout(_, _ int) (int, int) { return 600, 600 }
