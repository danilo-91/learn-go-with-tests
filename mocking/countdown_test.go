package main

import (
    "testing"
    "bytes"
)

func TestCountdown(t *testing.T) {
    b := &bytes.Buffer{}

    Countdown(b)

    got := b.String()
    want := `3
2
1
Go!
`

    if got != want {
        t.Errorf("wanted %q, but got %q", want, got)
    }
}
