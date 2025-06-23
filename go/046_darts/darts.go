package darts

import "math"

// Score calculates the points earned for a dart landing at coordinates (x, y)
func Score(x, y float64) int {
	// Calculate the distance from the center (0, 0)
	distance := math.Sqrt(x*x + y*y)

	if distance <= 1 {
		return 10
	} else if distance <= 5 {
		return 5
	} else if distance <= 10 {
		return 1
	} else {
		return 0
	}
}
