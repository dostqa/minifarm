package worldobjects

import (
	"minifarm/internal/runtime/territorygeneration/matrices"
)

func setRandomTree(mapData *matrices.Matrix) {
	tree := getRandomTreePoint((*mapData).N())

	for {
		if isMapPointEmpty((*mapData), tree.X, tree.Y) {
			(*mapData)[tree.X][tree.Y] = tree.TreeType
			break
		} else {
			tree = getRandomTreePoint((*mapData).N())
		}

	}
}

func GenerateTree(mapData *matrices.Matrix, n int) {
	for i := 1; i <= n; i++ {
		setRandomTree(mapData)
	}
	//fmt.Println("генерировать деревья кончил")
}
