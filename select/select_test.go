package main

import "testing"

func TestRacer(t *testing.T) {
    slowURL := "https://facebook.com"
    fastURL := "https://www.quii.dev"

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("wanted %q but got %q", want, got)
    }
}
