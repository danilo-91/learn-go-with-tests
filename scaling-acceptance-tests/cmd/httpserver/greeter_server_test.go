package main

import (
	"fmt"
	"github.com/isedaniel/go-specs-greet/adapters"
	"github.com/isedaniel/go-specs-greet/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = adapters.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecifications(t, driver)
}
