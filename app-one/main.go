package main

import (
	"log"
	"net/http"
)

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{map[string]int{}}
}
type MemoryStorage struct{
	storage map[string]int
}

func (m *MemoryStorage) GetPlayerScore(name string) int {
	return m.storage[name]
}

func (m *MemoryStorage) RecordAddScoreCall(name string) {
	m.storage[name]++
}

func main() {
	server := &PlayerServer{NewMemoryStorage()}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
