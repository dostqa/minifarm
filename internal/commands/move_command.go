package commands

import (
	"minifarm/internal/actors"
	"minifarm/internal/gametypes"
)

type MoveCommand struct {
	receiver  actors.Receiver
	direction gametypes.Vector
}

func (c MoveCommand) Execute() {
	if mover, ok := c.receiver.(actors.Mover); ok {
		mover.Move(c.direction)
	}
}

func (c MoveCommand) Undo() {} // Функция маркер

func NewMoveCommand(receiver actors.Receiver, direction gametypes.Vector) *MoveCommand {
	return &MoveCommand{
		receiver:  receiver,
		direction: direction,
	}
}

type StopCommand struct {
	receiver actors.Receiver
}

func (c StopCommand) Execute() {
	if mover, ok := c.receiver.(actors.Mover); ok {
		mover.Stop()
	}
}

func (c StopCommand) Undo() {} // Функция маркер

func NewStopCommand(receiver actors.Receiver) *StopCommand {
	return &StopCommand{
		receiver: receiver,
	}
}
