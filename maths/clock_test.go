package clock

import (
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	d := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(d)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
