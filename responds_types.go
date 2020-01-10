package mises

// indexResponse is returned by "/" on each service
type indexResponse struct {
	Name string `json:"name"`
}

// aliveResponse is returned by "/alive" on each service
type aliveResponse struct {
	Alive bool `json:"alive"`
}

// catchAllResponse is returned by "/alive" on each service
type catchAllResponse struct {
	Error string `json:"error"`
}
