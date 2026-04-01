package physic

import (
	"fmt"
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
)

var (
	DefaultHandler Handler
)

type Handler struct {
}

func (h Handler) Handle(event events.Event) {
	switch event.Type() {
	case gametypes.ToolUsedEventType:
		fmt.Println("Event was got: ", event)
	}

}

func NewHandler() *Handler {
	return &Handler{}
}
