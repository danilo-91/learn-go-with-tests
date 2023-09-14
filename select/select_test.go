package main

import (
    "testing"
    "time"
    "net/http"
    "net/http/httptest"
)

func TestRacer(t *testing.T) {
    slowServer := delayedServer(20 * time.Millisecond)
    fastServer := delayedServer(0 * time.Millisecond)

    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("wanted %q but got %q", want, got)
    }
}

func delayedServer(d time.Duration) *httptest.Server {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(d)
        w.WriteHeader(http.StatusOK)
    }))
    return server
}
