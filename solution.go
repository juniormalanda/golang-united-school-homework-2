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

type figureSidesType uint

const (
	SidesCircle figureSidesType = iota
	_
	_
	SidesTriangle figureSidesType = iota
	SidesSquare   figureSidesType = iota
)

func CalcSquare(sideLen float64, sidesNum figureSidesType) float64 {
	switch sideLen {
	case 0:
		return math.Pi * math.Pow(sideLen, 2)
	case 3:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
