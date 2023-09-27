package clock

import (
	"math"
	"time"
)

// Represents a two dimensional coordinate
type Point struct {
	X float64
	Y float64
}

// Returns the top point {X Y} of the second hand of an analogue clock
func SecondHandPoint(t time.Time) Point {
	angle := SecToRadian(t)
	X := math.Sin(angle)
	Y := math.Cos(angle)
	return Point{X, Y}
}

// Convert a time.Time.Second() to radians
func SecToRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
