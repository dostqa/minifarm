package actors

import "minifarm/internal/gametypes"

type Receiver any

type Mover interface {
	Move(gametypes.Vector)
}

type Tooler interface {
	UseActiveTool()
	ChangeActiveTool(int)
}
