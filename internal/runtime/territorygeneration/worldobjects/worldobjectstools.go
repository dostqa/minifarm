package worldobjects

import (
	"minifarm/internal/runtime/territorygeneration/matrices"
	"minifarm/internal/runtime/territorygeneration/randomness"
)

// isUsed проверяет не использовалась ли точка до этого.
func isUsed(points *[]RiverPoint, checkingP RiverPoint) bool {
	for _, p := range *points {
		if p.X == checkingP.X && p.Y == checkingP.Y {
			return true
		}
	}
	return false
}

func isMapPointEmpty(mapData matrices.Matrix, x, y int) bool {
	if mapData[x][y] == 0 {
		return true
	} else {
		return false
	}
}

// isntOut проверяет не вышли ли мы за границу карты
func isntOut(mapData *matrices.Matrix, checkingP RiverPoint) bool {
	if checkingP.X < 0 || checkingP.Y < 0 || checkingP.X >= (*mapData).N() || checkingP.Y >= (*mapData).M() {
		return false
	}
	return true
}

// getRandomCumNum возвращает случайное число,
// учитывая веса
func getRandomCumNum(max int) int {
	seq := randomness.NewSequenceSliceOneStart(max)
	symAbs := randomness.NewSymAbsSlice(max)
	cumul := randomness.NewCumulativeSlice(symAbs)

	r := randomness.GetRandomNum(cumul[len(cumul)-1])
	index := randomness.FindIndexByCumSum(cumul, r)

	return seq[index]
}

// nameToTile переводчик с человеческого на язык компьютера
var nameToTile = map[string]int{
	"empty":     0,
	"err":       1,
	"sourceU":   11,
	"sourceL":   12,
	"sourceD":   13,
	"sourceR":   14,
	"rbedH":     15,
	"rbedV":     16,
	"rcornerUL": 17,
	"rcornerDL": 18,
	"rcornerDR": 19,
	"rcornerUR": 20,
	"oneTree":   21,
	"twoTrees":  22,
}

// tileToName переводчик с компьютерного на язык человека
var tileToName = map[int]string{
	0:  "empty",
	1:  "err",
	11: "sourceU",
	12: "sourceL",
	13: "sourceD",
	14: "sourceR",
	15: "rbedH",
	16: "rbedV",
	17: "rcornerUL",
	18: "rcornerDL",
	19: "rcornerDR",
	20: "rcornerUR",
	21: "oneTree",
	22: "twoTrees",
}
