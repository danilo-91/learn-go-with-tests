package main_test

import (
	"go-with-tests/app-one"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Danilo score", func(t *testing.T) {
		req := newScoreRequest("Danilo")
		resp := httptest.NewRecorder()

		main.PlayerServer(resp, req)
		got := resp.Body.String()
		want := "20"
		assertString(t, got, want)
	})

	t.Run("returns Gabo score", func(t *testing.T) {
		req := newScoreRequest("Gabo")
		resp := httptest.NewRecorder()

		main.PlayerServer(resp, req)
		got := resp.Body.String()
		want := "30"
		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, but got %q", want, got)
	}
}

func newScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/" + name, nil)
	return req
}
	
