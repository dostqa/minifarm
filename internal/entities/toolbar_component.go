package entities

import (
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
)

// Интерфейс Tool описывает, что должен уметь инструмент
type Tool interface {
	Type() gametypes.ToolType
}

// ToolbarComponent - компонент панели инструментов.
// Любая сущность с этим компонентом может использовать инструменты
type ToolbarComponent struct {
	Publisher
	*PositionComponent
	tools  [5]Tool
	active int
}

func (bar *ToolbarComponent) UseActiveTool() {
	bar.Publish(events.NewToolUsedEvent(bar.Position(), bar.tools[bar.active].Type()))
}

func (bar *ToolbarComponent) ChangeActiveTool(i int) {
	// Логика, при которой происходит смена активного инструмента
	if (i >= 0) && (i <= 4) {
		if bar.active == i {
			return
		} else {
			// Непосредственно меняем активный слот
			bar.active = i

			// Создаём событие о смене активного слота
			bar.Publish(events.NewActiveToolChanged(bar.tools[bar.active].Type()))
		}
	} else {
		return
	}
}
