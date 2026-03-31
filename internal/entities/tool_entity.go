package entities

import "minifarm/internal/gametypes"

type Axe struct{}

func (a *Axe) Type() gametypes.ToolType {
	return gametypes.AxeTool
}
