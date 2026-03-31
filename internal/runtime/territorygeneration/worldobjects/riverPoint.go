package worldobjects

import "minifarm/internal/runtime/territorygeneration/worldtypes"

type RiverPoint struct {
	Point
	RiverType worldtypes.WorldObjectType
	direction Trend
	rotate    Turn
}

// NewRiverPoint возвращает новую точку реки;
// x, y - координаты точки на игровой карте;
// PointType - тип точки;
// RiverType - вид реки на карте;
// direction - в каком направлении будет течь река дальше относительно точки .
func NewRiverPoint(x, y int, RiverType worldtypes.WorldObjectType, direction Trend) RiverPoint {
	return RiverPoint{Point{x, y, "River"}, RiverType, direction, RelNone}
}

// Trend псевдоним для int;
// Описывает направление для следующего русла реки.
type Trend int

const (
	AbsNone Trend = iota
	AbsUp
	AbsRight
	AbsDown
	AbsLeft
)

type Turn int

// Turn псевдоним для int;
// Описывает поворот следующего русла реки относительно предыдущего.
const (
	RelStraight Turn = iota
	RelRight
	RelLeft
	RelNone
)
