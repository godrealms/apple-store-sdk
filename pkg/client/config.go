package client

import (
	"fmt"
	"strings"
	"time"
)

const (
	// DefaultTimeout is the default timeout for the HTTP client
	DefaultTimeout = 10 * time.Second

	// SandboxBaseURL is the base URL for the sandbox environment
	SandboxBaseURL = "https://api.storekit-sandbox.itunes.apple.com"

	// ProductionBaseURL is the base URL for the production environment
	ProductionBaseURL = "https://api.storekit.itunes.apple.com"

	// DefaultAPIVersion is the default version of the API
	DefaultAPIVersion = "v2"

	// DefaultRegion is the default region for API requests
	DefaultRegion = "CN"
)

// Config holds the configuration for the API client
type Config struct {
	BaseURL    string        // Base URL for the API
	Timeout    time.Duration // Timeout for HTTP requests
	Sandbox    bool          // Indicates whether to use sandbox mode
	APIVersion string        // API version (e.g., "v1", "v2")
	Region     string        // Region for API requests (e.g., "US", "CN")
	Iss        string        // Your issuer ID from the Keys page in App Store AppStoreConnectAPI (Ex: “57246542-96fe-1a63-e053-0824d011072a")
	Bid        string        // Your app’s bundle ID (Ex: “com.example.testbundleid”)
	Kid        string        // Key ID Your private key ID from App Store AppStoreConnectAPI (Ex: 2X9R4HXF34)
	//-----BEGIN PRIVATE KEY-----
	// Private key string corresponding to the private key ID from App Store AppStoreConnectAPI
	//-----END PRIVATE KEY-----
	PrivateKey string // Private key string corresponding to the private key ID from App Store AppStoreConnectAPI
	DebugMode  bool   // Debug mode for detailed logging (default: false)
}

// NewConfig creates a new configuration instance
func NewConfig(sandboxMode bool, Iss, Bid, Kid, privateKey string, timeout ...time.Duration) (*Config, error) {
	// Determine the BaseURL based on sandbox mode
	baseURL := ProductionBaseURL
	if sandboxMode {
		baseURL = SandboxBaseURL
	}

	// If no timeout is provided, use the default timeout
	var resolvedTimeout time.Duration
	if len(timeout) > 0 {
		resolvedTimeout = timeout[0]
	} else {
		resolvedTimeout = DefaultTimeout
	}

	if !strings.HasPrefix(strings.TrimSpace(privateKey), "-----BEGIN PRIVATE KEY-----") {
		privateKey = fmt.Sprintf(`-----BEGIN PRIVATE KEY-----
%s`, strings.TrimSpace(privateKey))
	}
	if !strings.HasSuffix(strings.TrimSpace(privateKey), "-----END PRIVATE KEY-----") {
		privateKey = fmt.Sprintf(`%s
-----END PRIVATE KEY-----`, strings.TrimSpace(privateKey))
	}

	config := &Config{
		BaseURL:    baseURL,
		Timeout:    resolvedTimeout,
		Sandbox:    sandboxMode,
		APIVersion: DefaultAPIVersion, // Default API version
		Region:     DefaultRegion,     // Default region
		Bid:        Bid,
		Iss:        Iss,
		Kid:        Kid,
		PrivateKey: privateKey,
	}

	return config, config.Validate()
}

// SetBaseUrl Set request address
func (c *Config) SetBaseUrl(url string) {
	c.BaseURL = url
}

// SetAPIVersion allows setting a custom API version
func (c *Config) SetAPIVersion(version string) *Config {
	c.APIVersion = version
	return c
}

// SetRegion allows setting a custom region
func (c *Config) SetRegion(region string) *Config {
	c.Region = region
	return c
}

// SetDebugMode allows enabling or disabling debug mode
func (c *Config) SetDebugMode(debug bool) *Config {
	c.DebugMode = debug
	return c
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// if c.APIKey == "" {
	// 	return fmt.Errorf("API key is missing")
	// }
	if c.BaseURL == "" {
		return fmt.Errorf("base URL is missing")
	}
	if c.Iss == "" {
		return fmt.Errorf("issuer ID is missing")
	}
	if c.Kid == "" {
		return fmt.Errorf("key ID is missing")
	}
	if c.APIVersion == "" {
		return fmt.Errorf("API version is missing")
	}
	if c.Region == "" {
		return fmt.Errorf("region is missing")
	}
	return nil
}
