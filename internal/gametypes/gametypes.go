package gametypes

type Point [2]float64

type Vector [2]float64

var (
	UpVector    = Vector{0, -1}
	RightVector = Vector{1, 0}
	DownVector  = Vector{0, 1}
	LeftVector  = Vector{-1, 0}
)

type ToolType int

const (
	NoTool ToolType = iota
	Pickaxe
	AxeTool
	SwordTool
	HoeTool
)

type EventType int

const (
	ToolUsedEventType EventType = iota
	ActiveToolChangedEventType
)
