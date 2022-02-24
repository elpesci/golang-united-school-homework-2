package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type Sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	noSides    = 0
	threeSides = 3
	fourSides  = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case noSides:
		return math.Pi * findSquare(sideLen)
	case threeSides:
		return findSquare(sideLen) * (math.Sqrt(3) / 4)
	case fourSides:
		return findSquare(sideLen)
	default:
		return 0
	}
}

func findSquare(x float64) float64 { return x * x }
