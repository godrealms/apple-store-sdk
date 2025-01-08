package client

import "net/http"

// Handler represents a function that processes an HTTP request and returns a response
type Handler func(req *http.Request) (*http.Response, error)

// Middleware represents a function that wraps a Handler with additional behavior
type Middleware func(next Handler) Handler
