package mises

import (
	"net/http"
)

// handleIndex is the http handler for root "/"
func (s *Server) handleIndex(serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Respond(w, r, indexResponse{Name: serviceName}, 200)
	}
}
