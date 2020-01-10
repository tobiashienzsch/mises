package mises_test

import (
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/tobiashienzsch/mises"
	"github.com/tobiashienzsch/mises/config"
)

func newTestServer() *httptest.Server {
	h := mises.NewServer(config.Config{SkipDatabase: true})
	h.AddRoutes("test_service", "/v1", func(*mux.Router) {})

	srv := httptest.NewServer(&h)
	return srv
}
