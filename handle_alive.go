package mises

import (
	"net/http"
)

// handleAlive returns a heartbeat by the server
func (s *Server) handleAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Respond(w, r, aliveResponse{Alive: true}, 200)
	}
}
