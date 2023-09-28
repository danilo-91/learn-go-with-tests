package clock

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSVGWriterMinuteHand(t *testing.T) {
    cases := []struct{
        time time.Time
        line Line
    }{
        {simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
    }

    for _, c := range cases {
        t.Run(timeName(c.time), func (t *testing.T) {
            b := bytes.Buffer{}
            SVGWriter(&b, c.time)

            svg := SVG{}
            xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Lines) {
                t.Errorf("wanted minute hand %+v, from SVG %+v", c.line, svg.Lines)
            }
        })
    }
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{150, 150 + 90}},
		{simpleTime(0, 0, 45), Point{150 - 90, 150}},
	}

	for _, c := range cases {
		t.Run(timeName(c.time), func(t *testing.T) {
			got := SecondHand(c.time)
			want := c.point
			assertPoint(t, got, want)
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(timeName(c.time), func(t *testing.T) {
			got := SecondHandPoint(c.time)
			want := c.point
			assertPoint(t, got, want)
		})
	}
}

func TestSecToRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, c := range cases {
		t.Run(timeName(c.time), func(t *testing.T) {
			want := c.angle
			got := SecToRadian(c.time)
			assertFloat64(t, got, want)
		})
	}
}

func TestSVGWriter(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}
	for _, c := range cases {
		t.Run(timeName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

            if !containsLine(c.line, svg.Lines) {
                t.Errorf("wanted %+v, from SVG %+v", c.line, svg.Lines)
            }
		})
	}
}

func assertPoint(t testing.TB, got, want Point) {
	t.Helper()
	if !approxEqualPoint(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func assertFloat64(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, but wanted %v", got, want)
	}
}

func simpleTime(hour, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hour, minutes, seconds, 0, time.UTC)
}

func timeName(t time.Time) string {
	return t.Format("15:04:05")
}

// Check that the difference between a and b is lesser than 1e-7
func approxEqual(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}

// Check that distance between Points a and b is lesser than 1e-7
func approxEqualPoint(a, b Point) bool {
	return approxEqual(a.X, b.X) &&
		approxEqual(a.Y, b.Y)
}

func containsLine(want Line, got []Line) bool {
    for _, line := range got {
		if line == want {
			return true
		}
	}
    return false
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}
