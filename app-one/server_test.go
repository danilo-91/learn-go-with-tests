package main_test

import (
	"go-with-tests/app-one"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Danilo score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Danilo", nil)
		resp := httptest.NewRecorder()

		main.PlayerServer(resp, req)
		got := resp.Body.String()
		want := "20"
		if got != want {
			t.Errorf("expected %q, but got %q", want, got)
		}
	})	
}
