package common

import (
	"math"
)

func CalculateDistance(x1, y1, x2, y2 int) float64 {
	diffInX := float64(x2 - x1)
	diffInY := float64(y2 - y1)

	distance := math.Sqrt(diffInX*diffInX + diffInY*diffInY)
	return math.Round(distance*100) / 100
}
