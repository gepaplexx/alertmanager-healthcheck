package webserver

import (
    "net/http"
    "fmt"
)

// Type for serving a simple message (probably for application readiness checks)
type RootEndpoint int

// Serves the message to a http endpoint
func (root RootEndpoint) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Alertmanager Health Check service!")
}

// Function for instantiating the RootEndpoint type
func NewRootEndpoint() RootEndpoint {
  var root RootEndpoint
  return root
}