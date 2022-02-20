package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type ShapeSides uint8

const (
	SidesCircle   ShapeSides = iota     // 0
	SidesTriangle ShapeSides = iota + 2 // 3
	SidesSquare   ShapeSides = iota + 2 // 4
)

type ShapeAreaCalc func(sideLen float64) float64

var areaCalcMap = map[ShapeSides]ShapeAreaCalc{
	SidesCircle:   circleAreaCalc,
	SidesTriangle: triangleAreaCalc,
	SidesSquare:   squareAreaCalc,
}

var sqrt3div4 = math.Sqrt(3.0) / 4.0

func squareAreaCalc(sideLen float64) float64 {
	return sideLen * sideLen
}

func triangleAreaCalc(sideLen float64) float64 {
	return sideLen * sideLen * sqrt3div4
}

func circleAreaCalc(sideLen float64) float64 {
	return math.Pi * sideLen * sideLen
}

func CalcSquare(sideLen float64, sidesNum ShapeSides) float64 {
	calc, found := areaCalcMap[sidesNum]
	if !found {
		return 0.0
	}

	return calc(sideLen)
}
