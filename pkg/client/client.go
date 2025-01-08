package client

import (
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/utils"
	"net/http"
	"time"
)

// Client is a wrapper for the HTTP client with middleware and configuration
type Client struct {
	httpHelper  *utils.HTTPHelper
	Config      *Config
	middlewares []Middleware
	logger      *utils.Logger // Unified logging interface
}

// NewClient creates a new API client with default middlewares
func NewClient(config *Config) *Client {
	// Validate the configuration
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	// Initialize HTTP client and logger
	httpHelper := utils.NewHTTPHelper(config.Timeout)
	logger := utils.NewLogger()

	// Create client instance
	client := &Client{
		httpHelper:  httpHelper,
		Config:      config,
		middlewares: []Middleware{},
		logger:      logger,
	}

	// Add default middlewares (logging and retry)
	client.Use(client.LoggingMiddleware())
	client.Use(client.RetryMiddleware(3, 2*time.Second)) // Default: 3 retries with 2s delay

	return client
}

// Use adds a middleware to the client
func (c *Client) Use(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

// LoggingMiddleware adds debug logs for requests and responses
func (c *Client) LoggingMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(req *http.Request) (*http.Response, error) {
			if c.Config.DebugMode {
				c.logger.Debug(fmt.Sprintf("Request URL: %s", req.URL.String()))
				c.logger.Debug(fmt.Sprintf("Request Method: %s", req.Method))
			}

			resp, err := next(req)

			if err != nil {
				if c.Config.DebugMode {
					c.logger.Debug(fmt.Sprintf("Request failed: %v", err))
				}
				return nil, err
			}

			if c.Config.DebugMode {
				c.logger.Debug(fmt.Sprintf("Response Status: %s", resp.Status))
			}

			return resp, nil
		}
	}
}

// Do sends an HTTP request with middleware support
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// Apply middlewares
	handler := func(req *http.Request) (*http.Response, error) {
		return c.httpHelper.Client.Do(req)
	}

	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	// Execute the request
	return handler(req)
}

func (c *Client) RetryMiddleware(maxRetries int, delay time.Duration) Middleware {
	return func(next Handler) Handler {
		return func(req *http.Request) (*http.Response, error) {
			var resp *http.Response
			var err error

			for attempt := 0; attempt <= maxRetries; attempt++ {
				// Call the next handler
				resp, err = next(req)

				// If no error or response is successful, return immediately
				if err == nil && resp.StatusCode < 500 {
					return resp, nil
				}

				// Log retry attempt
				c.logger.Warn("Retry attempt %d/%d failed: %v\n", attempt+1, maxRetries, err)

				// Wait before the next retry
				time.Sleep(delay)
			}

			// Return the last error after exhausting retries
			return nil, fmt.Errorf("all retries failed: %w", err)
		}
	}
}

// Get is a helper for GET requests
func (c *Client) Get(endpoint string, headers map[string]string, params map[string]string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s/%s/%s?%s", c.Config.BaseURL, c.Config.APIVersion, c.Config.Region, endpoint, utils.BuildQueryParams(params))
	return c.httpHelper.Get(url, headers)
}

// Post is a helper for POST requests
func (c *Client) Post(endpoint string, body []byte, headers map[string]string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.Config.BaseURL, c.Config.APIVersion, c.Config.Region, endpoint)
	return c.httpHelper.Post(url, body, headers)
}
