package handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gitpods/gitpods/handler"
	"github.com/gitpods/gitpods/store"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/gorilla/sessions"
	"github.com/pressly/chi"
	"github.com/stretchr/testify/assert"
)

var (
	httpTestClient = &http.Client{Timeout: 5 * time.Second}
)

func TestApiNotFound(t *testing.T) {
	r := DefaultTestAuthRouter()
	res, content, err := Request(r, http.MethodGet, "/404", nil)
	assert.NoError(t, err)
	assertNotFoundJson(t, res, content)
}

func TestApiNotFoundUnauthorized(t *testing.T) {
	r := DefaultTestRouter()
	res, content, err := Request(r, http.MethodGet, "/404", nil)
	assert.NoError(t, err)
	assertUnauthorized(t, res, content)
}

// Helpers

func DiscardMetrics() handler.RouterMetrics {
	return handler.RouterMetrics{
		LoginAttempts: discard.NewCounter(),
	}
}

func DefaultTestRouter() *chi.Mux {
	return DefaultTestRouterWithStore(DefaultRouterStore())
}

func DefaultTestRouterWithStore(store handler.RouterStore) *chi.Mux {
	return handler.NewRouter(log.NewNopLogger(), DiscardMetrics(), store)
}
func DefaultTestAuthRouter() *chi.Mux {
	return DefaultTestAuthRouterWithStore(DefaultRouterStore())
}

func DefaultTestAuthRouterWithStore(store handler.RouterStore) *chi.Mux {
	return handler.NewAuthRouter(log.NewNopLogger(), DiscardMetrics(), store)
}

func DefaultRouterStore() *handler.RouterStore {
	usersStore := store.NewUsersInMemory()
	return &handler.RouterStore{
		AuthorizeStore: usersStore,
		UsersStore:     usersStore,
		CookieStore:    sessions.NewCookieStore([]byte("secret")),
	}
}

func Request(r *chi.Mux, method string, url string, payload []byte) (*http.Response, []byte, error) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(method, ts.URL+url, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}

	res, err := httpTestClient.Do(req)

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	return res, content, nil
}

func assertUnauthorized(t *testing.T, res *http.Response, content []byte) {
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", res.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"message":"Unauthorized"}`, string(content))
}

func assertNotFoundJson(t *testing.T, res *http.Response, content []byte) {
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", res.Header.Get("Content-Type"))
	assert.Equal(t, string(handler.JsonNotFound), string(content))
}
