package main_test

import (
	"go-with-tests/app-one"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
	addScoreCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordAddScoreCall(name string) {
	s.addScoreCalls = append(s.addScoreCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Danilo": 20,
			"Gabo":   25,
		},
		nil,
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

func TestPOSTScores(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &main.PlayerServer{&store}

	t.Run("returns accepted / call addScore on valid POST", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/players/Danilo", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatusCode(t, resp.Code, http.StatusAccepted)

		if len(store.addScoreCalls) < 1 {
			t.Fatalf("got %d calls, but expected 1", len(store.addScoreCalls))
		}

		if store.addScoreCalls[0] != "Danilo" {
			t.Errorf("wrong stored player name, got %q, but expected %q", store.addScoreCalls[0], "Danilo")
		}
	})
}

func TestRecordingWinsAndRetrieving(t *testing.T) {
	store := main.MemoryStorage{}
	server := main.PlayerServer{&store}
	player := "Danilo"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newScoreRequest(player))
	assertStatusCode(t, resp.Code, http.StatusOK)
	assertString(t, resp.Body.String(), "3")
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
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/" + name, nil)
	return req
}
