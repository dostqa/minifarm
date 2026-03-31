package worldobjects

import (
	"math/rand"
	wt "minifarm/internal/runtime/territorygeneration/worldtypes"
)

func getRandomTreePoint(n int) TreePoint {
	rX, rY := getRandomTreeCords(n, n)

	return NewTreePoint(rX-1, rY-1, wt.OneTree)
}

func getRandomTreeCords(xMax, yMax int) (int, int) {
	x := rand.Intn(xMax) + 1
	y := rand.Intn(yMax) + 1

	return x, y
}
