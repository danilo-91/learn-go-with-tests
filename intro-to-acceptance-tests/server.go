package main

import (
	"log"
	"net/http"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
	"github.com/quii/go-graceful-shutdown/acceptancetests"
)

func main() {
    httpServer := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
    server := gracefulshutdown.NewServer(httpServer)

    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("did not shutdown gracefully, may have lost responses: %+v", err)
    }

    log.Println("Shuted down gracefully, all responses were sent")
}
