package main

import (
	"context"
    "errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("create server and get response", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		serv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		serv.ServeHTTP(res, req)
		got := res.Body.String()

		if got != data {
			t.Errorf("expected %q, but got %q", data, got)
		}

		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})

	t.Run("cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world!"
		store := &SpyStore{response: data}
		serv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := &SpyResponseWriter{}

		serv.ServeHTTP(res, req)

		if res.written {
			t.Error("store was not told to cancel")
		}
	})
}
