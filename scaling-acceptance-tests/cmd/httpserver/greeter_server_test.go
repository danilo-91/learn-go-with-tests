package main

import (
	"context"
	"testing"
	"time"
	"net/http"

	"github.com/alecthomas/assert/v2"
	go_specs_greet "github.com/isedaniel/go-specs-greet"
	"github.com/isedaniel/go-specs-greet/specifications"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{
		Timeout: 1* time.Second,
	}

	driver := go_specs_greet.Driver{
		BaseURL: "http://localhost:8080",
		Client: &client,
	}
	specifications.GreetSpecifications(t, driver)
}
