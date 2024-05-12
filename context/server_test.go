package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
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
	data := "hello world"
	store := &SpyStore{response: data}
	svr := Server(store)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(req.Context())
	// after 5 seconds call cancel
	time.AfterFunc(5*time.Millisecond, cancel)
	req = req.WithContext(cancellingCtx)

	res := httptest.NewRecorder()

	svr.ServeHTTP(res, req)

	if !store.cancelled {
		t.Error("store was not told to cancel")
	}
}
