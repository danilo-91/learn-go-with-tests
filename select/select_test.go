package main

import (
    "testing"
    "time"
    "net/http"
    "net/http/httptest"
)

func TestRacer(t *testing.T) {
    t.Run("return faster server", func (t *testing.T) {
        slowServer := delayedServer(20 * time.Millisecond)
        fastServer := delayedServer(0 * time.Millisecond)
        defer slowServer.Close()
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got, err := Racer(slowURL, fastURL)

        if err != nil {
            t.Errorf("not expected error %v", err)
        }

        if got != want {
            t.Errorf("wanted %q but got %q", want, got)
        }
    })

    t.Run("error after 10 seconds without response", func (t *testing.T) {
        s := delayedServer(25 * time.Millisecond)
        defer s.Close()

        _, err := ConfigurableRacer(s.URL, s.URL, 20*time.Millisecond)
        
        if err == nil {
            t.Errorf("expected error")
        }
    })
}

func delayedServer(d time.Duration) *httptest.Server {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(d)
        w.WriteHeader(http.StatusOK)
    }))
    return server
}
