package main

import (
    "fmt"
    "net/http"
)

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        store.Cancel()
        fmt.Fprintf(w, store.Fetch())
    }
}

type Store interface {
    Fetch() string
    Cancel()
}
