package client

import (
	"fmt"
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
	APIKey     string        // API Key for authentication
	BaseURL    string        // Base URL for the API
	Timeout    time.Duration // Timeout for HTTP requests
	Sandbox    bool          // Indicates whether to use sandbox mode
	APIVersion string        // API version (e.g., "v1", "v2")
	Region     string        // Region for API requests (e.g., "US", "CN")
	TeamID     string        // Apple Developer Team ID (for JWT authentication)
	BundleID   string        // Your application’s Bundle ID (Ex: “com.example.testbundleid”)
	KeyID      string        // Apple Developer Key ID (for JWT authentication)
	//-----BEGIN PRIVATE KEY-----
	// Private key string corresponding to the private key ID from App Store Connect
	//-----END PRIVATE KEY-----
	PrivateKey string // Private key string corresponding to the private key ID from App Store Connect
	DebugMode  bool   // Debug mode for detailed logging (default: false)
}

// NewConfig creates a new configuration instance
func NewConfig(sandboxMode bool, apiKey, teamID, bundleID, keyID, privateKey string, timeout ...time.Duration) *Config {
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

	return &Config{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		Timeout:    resolvedTimeout,
		Sandbox:    sandboxMode,
		APIVersion: DefaultAPIVersion, // Default API version
		Region:     DefaultRegion,     // Default region
		BundleID:   bundleID,
		TeamID:     teamID,
		KeyID:      keyID,
		PrivateKey: privateKey,
		DebugMode:  false, // Default to false
	}
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
	if c.APIKey == "" {
		return fmt.Errorf("API key is missing")
	}
	if c.BaseURL == "" {
		return fmt.Errorf("Base URL is missing")
	}
	if c.TeamID == "" {
		return fmt.Errorf("Team ID is missing")
	}
	if c.KeyID == "" {
		return fmt.Errorf("Key ID is missing")
	}
	if c.APIVersion == "" {
		return fmt.Errorf("API version is missing")
	}
	if c.Region == "" {
		return fmt.Errorf("Region is missing")
	}
	return nil
}
