package worldobjects

import (
	"minifarm/internal/runtime/territorygeneration/randomness"
	wt "minifarm/internal/runtime/territorygeneration/worldtypes"
)

// GetRandomPoint возвращает случайную точку на карте, с записанными координатами;
// pType созданной точки - "empty";
// trend - none;
// Предполагается, что карта квадратная,
// поэтому мы передаём только количество столбцов,
// чтобы случайная точка была создана в пределах карты.
func getRandomRiverPoint(n int) RiverPoint {
	rX, rY := getRandomRiverCords(n, n)

	return NewRiverPoint(rX-1, rY-1, wt.Grass, AbsNone)
}

func GetRandomRiverBed() int {
	var riverBedTypes = []int{0, 1, 2}
	var riverBedWeights = []int{1, 1, 4}
	cumul := randomness.NewCumulativeSlice(riverBedWeights)

	r := randomness.GetRandomNum(cumul[len(cumul)-1])
	index := randomness.FindIndexByCumSum(cumul, r)
	return riverBedTypes[index]
}

// GetRandomCords возвращает случайные координаты;
func getRandomRiverCords(xMax, yMax int) (int, int) {
	x := getRandomCumNum(xMax)
	y := getRandomCumNum(yMax)

	return x, y
}
