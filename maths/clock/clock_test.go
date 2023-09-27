package clock

import (
	"math"
	"testing"
    "bytes"
    "encoding/xml"
	"time"
)

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

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := simpleTime(0, 0, 0)
	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}
	t.Errorf("Expected second hand x2 = %+v and y2 = %+v from SVG o %v", x2, y2, b.String())
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

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line []struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}
