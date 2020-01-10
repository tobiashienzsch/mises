package mises

import (
	"net/http"
)

// HandleCatchAll handles all none specified routes. This needs to be
// added as the last route in setup of the individual service.
func (s *Server) HandleCatchAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Respond(w, r, catchAllResponse{Error: "endpoint not found"}, http.StatusNotFound)
	}
}
