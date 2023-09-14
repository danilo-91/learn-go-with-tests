package main

import (
    "testing"
    "fmt"
)

func TestWalk(t *testing.T) {

    expected := "Something else"
    var got []string

    x := struct {
        Name string
    }{expected}

    walk(x, func (input string) {
        got = append(got, input)
    })
    fmt.Println("got: ", got)

    if len(got) != 1 {
        t.Errorf("expected %d function calls, but got %d", 1, len(got))
    }

    if got[0] != expected {
        t.Errorf("expected %q, but gotten %q", expected, got[0])
    }
}
