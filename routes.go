package mises

import (
	"github.com/gorilla/mux"
)

// AddRoutes maps all common http endpoints & then calls the service routes function.
func (s *Server) AddRoutes(serviceName string, apiPrefix string, createServiceRoutes func(*mux.Router)) {
	api := s.Router.PathPrefix(apiPrefix).Subrouter()

	api.HandleFunc("/", s.handleIndex(serviceName)).Methods("GET")
	api.HandleFunc("/alive", s.handleAlive()).Methods("GET")

	createServiceRoutes(api)

	s.Router.PathPrefix("/").Handler(s.HandleCatchAll())
	s.Router.Use(s.LoggingMiddleware())

}
