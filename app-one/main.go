package main

import (
	"log"
	"net/http"
)

type MemoryStorage struct{}

func (m *MemoryStorage) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&MemoryStorage{}}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
