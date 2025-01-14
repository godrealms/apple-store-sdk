package client

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
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
func (c *Client) Get(endpoint string, headers map[string]string, params any) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s?%s", c.Config.BaseURL, endpoint, utils.BuildQueryParams(params))
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT()
	return c.httpHelper.Get(url, headers)
}

// Post is a helper for POST requests
func (c *Client) Post(endpoint string, body []byte, headers map[string]string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.Config.BaseURL, c.Config.APIVersion, c.Config.Region, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT()
	return c.httpHelper.Post(url, body, headers)
}

func (c *Client) PUT(endpoint string, headers map[string]string, body []byte) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT()
	return c.httpHelper.Put(url, body, headers)
}

func (c *Client) Patch(endpoint string, headers map[string]string, parameters any) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT()
	body, err := json.Marshal(parameters)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return c.httpHelper.Path(url, body, headers)
}

func (c *Client) Delete(endpoint string, headers map[string]string, body []byte) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s", c.Config.BaseURL, endpoint)
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Accept"] = "application/json"
	headers["Authorization"] = c.GenerateAuthorizationJWT()
	return c.httpHelper.Delete(url, body, headers)
}

// ParsePrivateKey Parse private keys in PEM format, supporting PKCS#8 and EC private keys
func ParsePrivateKey(pemKey string) (*ecdsa.PrivateKey, error) {
	// 解析 PEM 格式
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block: invalid format or empty key")
	}

	// 检查私钥类型
	var privateKey *ecdsa.PrivateKey
	var err error

	switch block.Type {
	case "PRIVATE KEY": // PKCS#8 格式
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, errors.New("failed to parse PKCS#8 private key: " + err.Error())
		}
		// 确保是 EC 私钥
		if ecKey, ok := key.(*ecdsa.PrivateKey); ok {
			privateKey = ecKey
		} else {
			return nil, errors.New("parsed key is not an ECDSA private key")
		}
	case "EC PRIVATE KEY": // EC 私钥格式
		privateKey, err = x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, errors.New("failed to parse EC private key: " + err.Error())
		}
	default:
		return nil, errors.New("unsupported private key type: " + block.Type)
	}

	return privateKey, nil
}

func (c *Client) GenerateAuthorizationJWT() string {
	privateKey, err := ParsePrivateKey(c.Config.PrivateKey)
	if err != nil {
		log.Printf("failed to parse private key: %v", err)
		return ""
	}
	// 创建 JWT 的 Header 和 Claims
	now := time.Now()
	claims := jwt.MapClaims{
		"iss": c.Config.TeamID,                  // Apple Team ID
		"iat": now.Unix(),                       // CURRENT TIMESTAMP
		"exp": now.Add(30 * time.Minute).Unix(), // Expiration time (30 minutes)
		"aud": "appstoreconnect-v1",             // Fixed value appstoreconnect-v1
		"bid": c.Config.BundleID,
	}
	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = c.Config.KeyID // Set Header's kid (key ID)
	// 使用私钥签名
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("failed to sign token")
		return ""
	}
	return fmt.Sprintf("Bearer %s", signedToken)
}
