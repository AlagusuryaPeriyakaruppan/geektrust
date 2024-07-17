package common

import (
	"fmt"
	"math"
	"strconv"
)

func ConvertStringToInt(value, errorMsg string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(errorMsg, err)
		return 0, err
	}
	return result, nil
}

func CalculateDistance(x1, y1, x2, y2 int) float64 {
	diffInX := float64(x2 - x1)
	diffInY := float64(y2 - y1)

	distance := math.Sqrt(diffInX*diffInX + diffInY*diffInY)
	return math.Round(distance*100) / 100
}
