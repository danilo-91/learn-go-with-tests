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

const (
	secHandLength = 90
	centreX       = 150
	centreY       = 150
)

// Returns the top point of the second hand of an analog clock ready for the SVG
func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secHandLength, p.Y * secHandLength} // scale
	p = Point{p.X, p.Y * -1}                            // flip (Y grows from top to bottom)
	p = Point{p.X + centreX, p.Y + centreY}             // translate to centre
	return p
}

// Returns the top point {X Y} of the second hand of an radius one analog clock
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
