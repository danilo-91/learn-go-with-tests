package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.String(), "/players/")

	if player == "Danilo" {
		fmt.Fprintf(w, "20")
	}
	if player == "Gabo" {
		fmt.Fprint(w, "30")
	}
}
