package main

import (
    "testing"
    "net/http/httptest"
    "net/http"
    "time"
    "context"
)

type SpyStore struct {
    response string
    cancelled bool
}

func (s *SpyStore) Fetch() string {
    time.Sleep(100 * time.Millisecond)
    return s.response
}

func (s *SpyStore) Cancel() {
    s.cancelled = true
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
        time.AfterFunc(5 * time.Millisecond, cancel)
        req = req.WithContext(cancellingCtx)

        res := httptest.NewRecorder()

        serv.ServeHTTP(res, req)

        if !store.cancelled {
            t.Error("store was not told to cancel")
        }
    })
}
