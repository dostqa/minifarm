package entities

import (
	"math"
	"minifarm/internal/gametypes"
)

type Ticker interface {
	Now() int
	TicksPerSecond() int
}

type StateComponent struct {
	Ticker

	current gametypes.StateName
	// В тиках
	startTime   int
	minDuration int
	maxDuration int
}

func (state *StateComponent) SetState(newState gametypes.StateName) {

	if state.Now()-state.startTime >= state.minDuration {

		switch newState {
		case gametypes.IdleStateName:
			state.current = gametypes.IdleStateName
			state.startTime = state.Now()
			state.minDuration = 0
			state.maxDuration = math.MaxInt

		case gametypes.MoveStateName:
			state.current = gametypes.MoveStateName
			state.startTime = state.Now()
			state.minDuration = 0
			state.maxDuration = math.MaxInt
		}
	}
}

func (state *StateComponent) GetState() gametypes.StateName {
	return state.current
}
