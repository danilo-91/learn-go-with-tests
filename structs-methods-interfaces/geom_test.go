package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	perimeterTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{3.0}, 2 * 3 * math.Pi},
	}

	for _, test := range perimeterTest {
		got := test.shape.Perimeter()
		checkFloat64(t, got, test.want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
        name string
		shape Shape
		want  float64
	}{
        {name: "Rectangle", shape: Rectangle{10.0, 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{3.0}, want: math.Pi * 3 * 3},
		{name: "Triangle", shape: Triangle{15.0, 6.0}, want: 45.0},
	}

	for _, test := range areaTests {
        t.Run(test.name, func (t *testing.T) {
            got := test.shape.Area()
            checkFloat64(t, got, test.want)
        })
	}
}

func checkFloat64(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("expected %g but got %g", want, got)
	}
}
