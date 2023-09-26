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

func simpleTime(hour, minutes, seconds int) time.Time {
    return time.Date(1337, time.January, 1, hour, minutes, seconds, 0, time.UTC)
}

func timeName(t time.Time) string {
    return t.Format("15:04:05")
}
