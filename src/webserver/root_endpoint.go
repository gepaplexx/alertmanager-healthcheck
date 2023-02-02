package webserver

import (
	"fmt"
	"net/http"
)

// RootEndpoint Type for serving a simple message (probably for application readiness checks)
type RootEndpoint int

// Serves the message to a http endpoint
func (root RootEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the Alertmanager Health Check service!")
	if err != nil {
		return
	}
}

// NewRootEndpoint Function for instantiating the RootEndpoint type
func NewRootEndpoint() RootEndpoint {
	var root RootEndpoint
	return root
}
