package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.String(), "/players/")
	fmt.Fprintf(w, playerScore(player))
}

func playerScore(name string) string {
	switch name {
	case "Danilo":
		return "20"
	case "Gabo":
		return "30"
	default:
		return ""
	}
}
