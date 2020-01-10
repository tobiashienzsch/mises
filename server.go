package mises

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/gorilla/mux"

	"github.com/tobiashienzsch/mises/config"
)

// Server holds all configuration or an http server
type Server struct {
	IP       string
	Port     int
	Router   *mux.Router
	Config   config.Config
	Database *pop.Connection
	Logger   *Logger
}

// NewServer returns a default configured http server
func NewServer(conf config.Config) Server {
	env := envy.Get("GO_ENV", "development")

	logger := NewLogger()
	logger.Debug = env == "development"
	logger.Color = true
	logger.Infof("GO_ENV: %s\n", env)

	var db *pop.Connection
	if !conf.SkipDatabase {
		pop.SetLogger(logger.createDatabaseLogger())
		database, err := CreateDatabaseConnection(env)
		if err != nil {
			logger.Fatal(err)
		}

		db = database
		logger.Infof("DATABASE: %s\n", db.Dialect.Name())
	}

	srv := Server{
		IP:       "0.0.0.0",
		Port:     3000,
		Config:   conf,
		Router:   mux.NewRouter(),
		Database: db,
		Logger:   logger,
	}

	return srv
}

// StartHTTPServer starts the server
func (s *Server) StartHTTPServer() {
	srvAddr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	s.Logger.Infof("START: %s\n", srvAddr)

	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		str, _ := route.GetPathTemplate()
		s.Logger.Info(str)
		return nil
	})

	httpSrv := &http.Server{
		Handler: s,
		Addr:    srvAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.Logger.Error(httpSrv.ListenAndServe())
}

// ServeHTTP dispatches the handler registered in the matched route.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Decode the json body from an http request
func (s *Server) Decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// Respond helper function for http/json responses
func (s *Server) Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.Logger.Errorf("Error encoding json data in response")
		}
	}
}
