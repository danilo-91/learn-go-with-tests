package main

import (
    "testing"
    "bytes"
)

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

func TestCountdown(t *testing.T) {
    b := &bytes.Buffer{}
    sl := &SpySleeper{}

    Countdown(b, sl)

    got := b.String()
    want := `3
2
1
Go!
`

    if got != want {
        t.Errorf("wanted %q, but got %q", want, got)
    }

    if sl.Calls != 3 {
        t.Errorf("sleep should be called 3 times! was called %d", sl.Calls)
    }
}
