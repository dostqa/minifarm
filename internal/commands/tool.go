package commands

import (
	"minifarm/internal/actors"
)

type UseActiveToolCommand struct {
	receiver actors.Receiver
}

func (c *UseActiveToolCommand) Execute() {
	if user, ok := c.receiver.(actors.Tooler); ok {
		user.UseActiveTool()
	}
}

func (c *UseActiveToolCommand) Undo() {} // Функция маркер

func NewUseActiveToolCommand(receiver actors.Receiver) *UseActiveToolCommand {
	return &UseActiveToolCommand{
		receiver: receiver,
	}
}

type ChangeActiveToolCommand struct {
	receiver actors.Receiver
	active   int
}

func (c *ChangeActiveToolCommand) Execute() {
	if tooler, ok := c.receiver.(actors.Tooler); ok {
		tooler.ChangeActiveTool(c.active)
	}
}

func (c *ChangeActiveToolCommand) Undo() {} // Функция маркер

func NewChangeActiveToolCommand(receiver actors.Receiver, active int) *ChangeActiveToolCommand {
	return &ChangeActiveToolCommand{
		receiver: receiver,
		active:   active,
	}
}
