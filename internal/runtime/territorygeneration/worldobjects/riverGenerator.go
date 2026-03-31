package worldobjects

import (
	"math/rand"
	"minifarm/internal/runtime/territorygeneration/matrices"
	wt "minifarm/internal/runtime/territorygeneration/worldtypes"
)

// setRiverSource создаёт на карте исток реки;
// А также записыват в срез использованных точек исток и место будущего русла
func setRiverSource(mapData *matrices.Matrix, usedPoints *[]RiverPoint) {
	source := getRandomRiverPoint((*mapData).N())
	for {
		if isUsed(usedPoints, source) {
			source = getRandomRiverPoint((*mapData).N())
		} else {
			break
		}

	}
	source = getRiverSourceDirection((*mapData).M(), (*mapData).N(), source.X, source.Y)
	var sourceType wt.WorldObjectType = source.RiverType

	var np RiverPoint
	switch sourceType {
	case wt.SourceU:
		np = NewRiverPoint(source.X, source.Y-1, wt.RbedV, AbsUp)
	case wt.SourceD:
		np = NewRiverPoint(source.X, source.Y+1, wt.RbedV, AbsDown)
	case wt.SourceR:
		np = NewRiverPoint(source.X+1, source.Y, wt.RbedH, AbsRight)
	case wt.SourceL:
		np = NewRiverPoint(source.X-1, source.Y, wt.RbedH, AbsLeft)
	}

	(*usedPoints) = append((*usedPoints), source, np)
}

// getRiverSourceDirection случайно определяет направление основываясь на координатах переданной точки,
// А также размерах карты;
// Русло всегда стремится быть направлено в центр карты.
func getRiverSourceDirection(m, n, pX, pY int) RiverPoint {
	midM := m / 2
	midN := n / 2
	posPoints := make([]RiverPoint, 2)

	if pX > midM {
		posPoints[0] = NewRiverPoint(pX, pY, wt.SourceL, AbsLeft)
	} else {
		posPoints[0] = NewRiverPoint(pX, pY, wt.SourceR, AbsRight)
	}

	if pY > midN {
		posPoints[1] = NewRiverPoint(pX, pY, wt.SourceU, AbsUp)
	} else {
		posPoints[1] = NewRiverPoint(pX, pY, wt.SourceD, AbsDown)
	}

	return posPoints[rand.Intn(2)]
}

// setRiverBed записывает точку русла реки на карту;
// mapData - игровая карта.
func setRiverBed(mapData *matrices.Matrix, usedPoints *[]RiverPoint, len *int) int {
	lp := getLastRiverPoint(usedPoints)
	np := nextRiverBedOnPoint(*usedPoints, lp)
	nnp := nextRiverBedOnPoint(*usedPoints, np)

	i := 0
	for {
		if isUsed(usedPoints, np) || isUsed(usedPoints, nnp) {
			np = nextRiverBedOnPoint(*usedPoints, lp)
			nnp = nextRiverBedOnPoint(*usedPoints, np)
			i += 1
		} else {
			if !isntOut(mapData, np) {
				return 1
			} else if !isntOut(mapData, nnp) {
				(*usedPoints) = append((*usedPoints), np)
				return 1
			}
			(*usedPoints) = append((*usedPoints), np, nnp)
			return 0
		}
		if i > 5 {
			if (*len) > 10 {
				(*len) = 0
				return 1
			}
			deleteLastRiverPoint(usedPoints)
			(*len) += 1
			return 0
		}
	}
}

// nextRiverBedOnPoint создаёт следующее русло реки, основываясь на последней точке.
func nextRiverBedOnPoint(usedPoints []RiverPoint, lp RiverPoint) RiverPoint {
	var np RiverPoint = createRiverBedOpts(usedPoints, lp)
	var direction Trend = lp.direction

	switch direction {
	case AbsUp:
		np.X = lp.X
		np.Y = lp.Y - 1
	case AbsDown:
		np.X = lp.X
		np.Y = lp.Y + 1
	case AbsRight:
		np.X = lp.X + 1
		np.Y = lp.Y
	case AbsLeft:
		np.X = lp.X - 1
		np.Y = lp.Y
	}

	return np
}

// createRiverBedOpts создаёт информацию о русле в поля Point, основываясь на последней точке.
func createRiverBedOpts(usedPoints []RiverPoint, lp RiverPoint) RiverPoint {

	var np RiverPoint
	var r int = GetNextRiverBedDirection(usedPoints)
	var direction Trend = lp.direction

	if direction == AbsDown {
		switch r {
		case 0:
			np.RiverType = wt.RcornerDR
			np.direction = AbsLeft
			np.rotate = RelRight
		case 1:
			np.RiverType = wt.RcornerDL
			np.direction = AbsRight
			np.rotate = RelLeft
		case 2:
			np.RiverType = wt.RbedV
			np.direction = AbsDown
			np.rotate = RelStraight
		}
	} else if direction == AbsUp {
		switch r {
		case 0:
			np.RiverType = wt.RcornerUL
			np.direction = AbsRight
			np.rotate = RelRight
		case 1:
			np.RiverType = wt.RcornerUR
			np.direction = AbsLeft
			np.rotate = RelLeft
		case 2:
			np.RiverType = wt.RbedV
			np.direction = AbsUp
			np.rotate = RelStraight
		}
	} else if direction == AbsRight {
		switch r {
		case 0:
			np.RiverType = wt.RcornerUR
			np.direction = AbsDown
			np.rotate = RelRight
		case 1:
			np.RiverType = wt.RcornerDR
			np.direction = AbsUp
			np.rotate = RelLeft
		case 2:
			np.RiverType = wt.RbedH
			np.direction = AbsRight
			np.rotate = RelStraight
		}
	} else if direction == AbsLeft {
		switch r {
		case 0:
			np.RiverType = wt.RcornerDL
			np.direction = AbsUp
			np.rotate = RelRight
		case 1:
			np.RiverType = wt.RcornerUL
			np.direction = AbsDown
			np.rotate = RelLeft
		case 2:
			np.RiverType = wt.RbedH
			np.direction = AbsLeft
			np.rotate = RelStraight
		}
	}
	return np
}

func GetNextRiverBedDirection(points []RiverPoint) int {
	var rightCount int
	var leftCount int
	var result int

	result = GetRandomRiverBed()

	if !(result == 2) {
		for _, p := range points {
			if p.rotate == RelRight {
				rightCount += 1
			} else if p.rotate == RelLeft {
				leftCount += 1
			}
		}

		if rightCount >= leftCount+2 {
			result = 1
			//fmt.Println("был когда права больше")
		} else if leftCount >= rightCount+2 {
			//fmt.Println("был когда лева больше")
			result = 0
		}
	}
	//fmt.Println("По итогу вернул: ", result)
	return result
}

// GenerateRiver создаёт полноценную реку
func GenerateRiver(mapData *matrices.Matrix, usedPoints *([]RiverPoint)) {
	var len int
	setRiverSource(mapData, usedPoints)

	var err int

	for {
		err = setRiverBed(mapData, usedPoints, &len)
		if err == 1 {
			break
		}
	}

	for _, p := range *usedPoints {
		(*mapData)[p.X][p.Y] = p.RiverType
	}

	//fmt.Println("РекуГенирировать кончил")
}

// getLastRiverPoint получает информацию о последней поставленной,
// или, в некоторых случаях, планируемой, клетки реки.
func getLastRiverPoint(usedPoints *[]RiverPoint) RiverPoint {
	return ((*usedPoints)[len((*(usedPoints)))-1])
}

func deleteLastRiverPoint(slice *[]RiverPoint) {
	if len(*slice) > 0 {
		*slice = (*slice)[:len(*slice)-1]
	}
}
