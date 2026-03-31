package events

import (
	"fmt"
	"minifarm/internal/gametypes"
)

type ToolUsedEvent struct {
	position gametypes.Point
	tool     gametypes.ToolType
}

func (e *ToolUsedEvent) String() string {
	return ("Применен активный инструмент на " + fmt.Sprintf("%#v", e.position))
}

func (e *ToolUsedEvent) Type() gametypes.EventType {
	return gametypes.ToolUsedEventType
}

func NewToolUsedEvent(pos gametypes.Point, tool gametypes.ToolType) *ToolUsedEvent {
	return &ToolUsedEvent{
		position: pos,
		tool:     tool,
	}
}

type ActiveToolChangedEvent struct {
	tool gametypes.ToolType
}

func (e *ActiveToolChangedEvent) String() string {
	return ("Активный инструмент изменен на " + fmt.Sprintf("%d", e.tool))
}

func (e *ActiveToolChangedEvent) Type() gametypes.EventType {
	return gametypes.ActiveToolChangedEventType
}

func NewActiveToolChanged(tool gametypes.ToolType) *ActiveToolChangedEvent {
	return &ActiveToolChangedEvent{
		tool: tool,
	}
}
