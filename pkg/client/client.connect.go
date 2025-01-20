package client

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/JWT"
	"github.com/godrealms/apple-store-sdk/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

// ConnectClient is a wrapper for the HTTP client with middleware and configuration
type ConnectClient struct {
	httpHelper  *utils.HTTPHelper
	Config      *Config
	middlewares []Middleware
	logger      *utils.Logger // Unified logging interface
}

// NewConnectClient creates a new API client with default middlewares
func NewConnectClient(config *Config) *ConnectClient {
	// Validate the configuration
	if err := config.Validate(); err != nil {
		panic(fmt.Sprintf("Invalid configuration: %v", err))
	}

	// Initialize HTTP client and logger
	httpHelper := utils.NewHTTPHelper(config.Timeout)
	logger := utils.NewLogger()

	// Create client instance
	client := &ConnectClient{
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
func (c *ConnectClient) Use(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

// LoggingMiddleware adds debug logs for requests and responses
func (c *ConnectClient) LoggingMiddleware() Middleware {
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
func (c *ConnectClient) Do(req *http.Request) (*http.Response, error) {
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

func (c *ConnectClient) RetryMiddleware(maxRetries int, delay time.Duration) Middleware {
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
func (c *ConnectClient) Get(endpoint string, headers map[string]string, params any) ([]byte, int, error) {
	endpoint = fmt.Sprintf("%s%s", endpoint, utils.BuildQueryParams(params))
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT("GET", endpoint)
	return c.httpHelper.Get(url, headers)
}

// Post is a helper for POST requests
func (c *ConnectClient) Post(endpoint string, body []byte, headers map[string]string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT("POST", endpoint)
	return c.httpHelper.Post(url, body, headers)
}

func (c *ConnectClient) PUT(endpoint string, headers map[string]string, body []byte) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT("PUT", endpoint)
	return c.httpHelper.Put(url, body, headers)
}

func (c *ConnectClient) Patch(endpoint string, headers map[string]string, parameters any) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT("PATCH", endpoint)
	body, err := json.Marshal(parameters)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return c.httpHelper.Path(url, body, headers)
}

func (c *ConnectClient) Delete(endpoint string, headers map[string]string, body []byte) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT("DELETE", endpoint)
	return c.httpHelper.Delete(url, body, headers)
}

func (c *ConnectClient) GenerateAuthorizationJWT(method string, endpoint string) string {
	privateKey, err := JWT.ParsePrivateKey(c.Config.PrivateKey)
	if err != nil {
		log.Printf("failed to parse private key: %v", err)
		return ""
	}
	// 创建 JWT 的 Header 和 Claims
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": c.Config.Iss,                     // Apple Team ID
		"iat": now.Unix(),                       // CURRENT TIMESTAMP
		"exp": now.Add(30 * time.Minute).Unix(), // Expiration time (30 minutes)
		"aud": "appstoreconnect-v1",             // Fixed value appstoreconnect-v1
		"scope": []string{
			fmt.Sprintf("%s %s", method, endpoint),
		},
	}
	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = c.Config.Kid // Set Header's kid (key ID)

	header, _ := json.Marshal(token.Header)
	log.Println("generated JWT Header:", string(header))
	payload, _ := json.Marshal(claims)
	log.Println("generated JWT Claims:", string(payload))

	// 使用私钥签名
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("failed to sign token")
		return ""
	}
	return fmt.Sprintf("Bearer %s", signedToken)
}
