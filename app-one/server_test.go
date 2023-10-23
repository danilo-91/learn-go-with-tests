package main_test

import (
	"go-with-tests/app-one"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Danilo": 20,
			"Gabo": 25,
		},
	}
	server := &main.PlayerServer{&store}

	t.Run("returns Danilo score", func(t *testing.T) {
		req := newScoreRequest("Danilo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)
		got := resp.Body.String()
		want := "20"
		assertString(t, got, want)
		assertStatusCode(t, resp.Code, http.StatusOK)
	})

	t.Run("returns Gabo score", func(t *testing.T) {
		req := newScoreRequest("Gabo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)
		got := resp.Body.String()
		want := "25"
		assertString(t, got, want)
		assertStatusCode(t, resp.Code, http.StatusOK)
	})

	t.Run("returns 404 err on unexisting player", func(t *testing.T) {
		req := newScoreRequest("Thompson")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := resp.Code
		want := http.StatusNotFound

		assertStatusCode(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q, but got %q", want, got)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, but wanted %d", got, want)
	}
}

func newScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/" + name, nil)
	return req
}
	
