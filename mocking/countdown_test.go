package main

import (
    "testing"
    "bytes"
)

func TestCountdown(t *testing.T) {
    b := &bytes.Buffer{}

    Countdown(b)

    got := b.String()
    want := "1\n2\n3\nGo!"

    if got != want {
        t.Errorf("wanted %q, but got %q", want, got)
    }
}
