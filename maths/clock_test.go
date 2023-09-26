package clock

import (
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	t.Run("midnight", func(t *testing.T) {
		hour := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
		want := Point{X: 150, Y: 150 - 90}
		got := SecondHand(hour)
		assertPoints(t, got, want)
	})

	t.Run("6 o'clock", func(t *testing.T) {
		hour := time.Date(1337, time.January, 1, 6, 0, 0, 0, time.UTC)
		want := Point{X: 150, Y: 150 + 90}
		got := SecondHand(hour)
		assertPoints(t, got, want)
	})
}

func TestSecToRadian(t *testing.T) {
	thirtySec := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := SecToRadian(thirtySec)
	assertFloat64(t, got, want)
}

func assertPoints(t testing.TB, got, want Point) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func assertFloat64(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, but wanted %v", got, want)
	}
}
