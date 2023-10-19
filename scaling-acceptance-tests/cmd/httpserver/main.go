package main

import (
	"github.com/isedaniel/go-specs-greet/adapters"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(adapters.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
