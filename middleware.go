package mises

import (
	"net/http"
)

// LoggingMiddleware logs every incoming request
func (s *Server) LoggingMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.Logger.Infof("%s: %s\n", r.Method, r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
