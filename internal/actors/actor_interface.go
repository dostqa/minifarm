package actors

import "minifarm/internal/gametypes"

// Receiver - актёр, получающий команды
type Receiver any

// Mover - актёр, способный двигаться
type Mover interface {
	Move(gametypes.Vector)
	Stop()
}

// Tooler - актёр, способный выполнить команды, связанные с инструментами
type Tooler interface {
	UseActiveTool()
	ChangeActiveTool(int)
}
