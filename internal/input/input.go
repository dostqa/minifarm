package input

import (
	"fmt"
	"minifarm/internal/actors"
	"minifarm/internal/commands"
	"minifarm/internal/gametypes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	DefaultInput Input
)

// Invoker вызывает команды, созданные во время
// работы Input
type Invoker interface {
	SetCommand(commands.Command)
}

// Структура Input контроллирует пользовательский ввод
type Input struct {
	invoker Invoker
}

func (i *Input) HandleInput(receiver actors.Receiver) error {
	i.invoker.SetCommand(commands.NewStopCommand(receiver))

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		i.invoker.SetCommand(commands.NewMoveCommand(receiver, gametypes.UpVector))
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		i.invoker.SetCommand(commands.NewMoveCommand(receiver, gametypes.RightVector))
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		i.invoker.SetCommand(commands.NewMoveCommand(receiver, gametypes.DownVector))
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		i.invoker.SetCommand(commands.NewMoveCommand(receiver, gametypes.LeftVector))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		i.invoker.SetCommand(commands.NewUseActiveToolCommand(receiver))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) {
		i.invoker.SetCommand(commands.NewChangeActiveToolCommand(receiver, 0))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) {
		i.invoker.SetCommand(commands.NewChangeActiveToolCommand(receiver, 1))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit3) {
		i.invoker.SetCommand(commands.NewChangeActiveToolCommand(receiver, 2))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit4) {
		i.invoker.SetCommand(commands.NewChangeActiveToolCommand(receiver, 3))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit5) {
		i.invoker.SetCommand(commands.NewChangeActiveToolCommand(receiver, 4))
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("exit")
	}
	return nil
}

func (i *Input) ConnectToInvoker(invoker Invoker) {
	i.invoker = invoker
}
