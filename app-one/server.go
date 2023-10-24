package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordAddScoreCall(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, r)
	case http.MethodPost:
		p.setScore(w, r)
	}
}

func (p *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {
	player := playerName(r)
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) setScore(w http.ResponseWriter, r *http.Request) {
	player := playerName(r)
	p.Store.RecordAddScoreCall(player)
	w.WriteHeader(http.StatusAccepted)
}

func playerName(r *http.Request) string {
	return strings.TrimPrefix(r.URL.String(), "/players/")
}
