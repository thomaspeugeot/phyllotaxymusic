package models

import "math"

// Convert degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Convert radians to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
