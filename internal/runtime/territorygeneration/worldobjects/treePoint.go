package worldobjects

import "minifarm/internal/runtime/territorygeneration/worldtypes"

type TreePoint struct {
	Point
	TreeType worldtypes.WorldObjectType
}

func NewTreePoint(x, y int, TreeType worldtypes.WorldObjectType) TreePoint {
	return TreePoint{Point{x, y, "Tree"}, TreeType}
}
