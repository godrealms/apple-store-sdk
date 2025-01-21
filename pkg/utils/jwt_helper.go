package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateAuthorizationJWT generates a JWT for App Store AppStoreServerAPI API
func GenerateAuthorizationJWT(IssuerID, BundleID, KeyID, PrivateKey string) (string, error) {
	// Read the private key from the .p8 file
	// privateKeyBytes, err := ioutil.ReadFile(PrivateKey)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to read private key file: %w", err)
	// }
	privateKeyBytes := []byte(PrivateKey)
	// Parse the private key
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	ecPrivateKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", errors.New("private key is not a valid ECDSA key")
	}

	// Create the JWT claims
	claims := jwt.MapClaims{
		"iss": IssuerID,                                // Issuer ID
		"iat": time.Now().Unix(),                       // Issued at
		"exp": time.Now().Add(30 * time.Minute).Unix(), // Expiration time (max 30 minutes)
		"aud": "appstoreconnect-v1",                    // Audience
		"bid": BundleID,
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Set the Key ID in the header
	token.Header["kid"] = KeyID

	// Sign the token with the private key
	signedToken, err := token.SignedString(ecPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}
