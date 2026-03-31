package physic

import (
	"fmt"
	"minifarm/internal/events"
)

var (
	DefaultHandler Handler
)

type Handler struct {
}

func (h Handler) Handle(event events.Event) {
	fmt.Println("Event was got: ", event)
}

func NewHandler() *Handler {
	return &Handler{}
}
