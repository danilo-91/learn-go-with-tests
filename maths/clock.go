package clock

import "time"

// Returns the second point of the hour hand of an analogue clock
func SecondHand(t time.Time) Point {
	return Point{}
}

// Represents a two dimensional coordinate
type Point struct {
	X float64
	Y float64
}
