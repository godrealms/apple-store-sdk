package client

import "errors"

// Define common client errors
var (
	ErrMissingAPIKey   = errors.New("missing API key")
	ErrInvalidResponse = errors.New("invalid response from server")
	ErrRequestFailed   = errors.New("request failed")
)
