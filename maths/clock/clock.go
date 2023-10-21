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

// Returns the top point of the second hand of an analog clock ready for the SVG
func SecondHand(t time.Time) Point {
	handLength := 90.0
	p := SecondHandPoint(t)
	return makeHand(p, handLength)
}

// Convert t time.Time to rads, then to Point{Sin(X), Cos(Y)} for Second Hand
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecToRadian(t))
}

// Convert a time.Time.Second() to radians
func SecToRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

// Returns the top point of the minute hand for an analog clock ready for the SVG
func MinuteHand(t time.Time) Point {
	handLength := 80.0
	p := MinuteHandPoint(t)
	return makeHand(p, handLength)
}

// Convert t time.Time to rads, then to Point{Sin(X), Cos(Y)} for Minute Hand
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinToRadian(t))
}

func MinToRadian(t time.Time) float64 {
	return (SecToRadian(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func HourHand(t time.Time) Point {
    handLength := 50.0
    p := HourHandPoint(t)
    return makeHand(p, handLength)
}

func HourHandPoint(t time.Time) Point {
    return angleToPoint(HourToRadian(t))
}

func HourToRadian(t time.Time) float64 {
	return (MinToRadian(t) / 12) +
		(math.Pi / (6 / float64(t.Hour() % 12)))
}

// Convert rad angle to Point{Sin(angle), Cos(angle)}
func angleToPoint(angle float64) Point {
	X := math.Sin(angle)
	Y := math.Cos(angle)
	return Point{X, Y}
}

const (
	centre = 150
)

// Receive Point{X, Y} and h handLength to return scaled, fliped and translated hand point
func makeHand(p Point, h float64) Point {
	p = Point{p.X * h, p.Y * h}           // scale
	p = Point{p.X, p.Y * -1}              // flip
	p = Point{p.X + centre, p.Y + centre} // translate to centre
	return p
}
