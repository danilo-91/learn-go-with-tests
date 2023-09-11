package main

import (
    "testing"
    "bytes"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Danilo")

    got := buffer.String()
    want := "Hello, Danilo"

    if got != want {
        t.Errorf("wanted %q, but got %q", want, got)
    }
}
