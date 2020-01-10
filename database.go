package mises

import (
	"github.com/gobuffalo/pop"
)

// CreateDatabaseConnection returns a pop connection
func CreateDatabaseConnection(env string) (*pop.Connection, error) {
	// Connect to db
	db, err := pop.Connect(env)
	if err != nil {
		return nil, err
	}

	// Enable database debug messages if in debug mode
	pop.Debug = env == "development"

	return db, nil
}
