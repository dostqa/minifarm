package gametypes

type Point [2]float64

type Vector [2]float64

var (
	ZeroVector  = Vector{0, 0}
	UpVector    = Vector{0, -1}
	RightVector = Vector{1, 0}
	DownVector  = Vector{0, 1}
	LeftVector  = Vector{-1, 0}
)

type ToolType int

const (
	NoTool ToolType = iota
	PickaxeTool
	AxeTool
	SwordTool
	HoeTool
)

type EventType int

const (
	// Связанные с инструментами
	ToolUsedEventType EventType = iota
	ActiveToolChangedEventType

	// Связанные со звуками
	EntityMovedEventType
)
