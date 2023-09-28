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

// Convert t time.Time to rads, then to Point{Sin(X), Cos(Y)} for Second Hand
func SecondHandPoint(t time.Time) Point {
    return angleToPoint(SecToRadian(t))
}

// Convert a time.Time.Second() to radians
func SecToRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

const (
    minHandLength = 80
)

// Returns the top point of the minute hand for an analog clock ready for the SVG
func MinuteHand(t time.Time) Point {
    p := MinuteHandPoint(t)
    p = Point{p.X*minHandLength, p.Y*minHandLength} // scale
    p = Point{p.X, p.Y * -1} // flip
    p = Point{p.X + centreX, p.Y + centreY} // translate to centre
    return p
}

// Convert t time.Time to rads, then to Point{Sin(X), Cos(Y)} for Minute Hand
func MinuteHandPoint(t time.Time) Point {
    return angleToPoint(MinToRadian(t))
}

func MinToRadian(t time.Time) float64 {
	return (SecToRadian(t) / 60) +
    (math.Pi / (30 / float64(t.Minute())))
}

// Convert rad angle to Point{Sin(angle), Cos(angle)}
func angleToPoint(angle float64) Point {
    X := math.Sin(angle)
    Y := math.Cos(angle)
    return Point{X, Y}
}
