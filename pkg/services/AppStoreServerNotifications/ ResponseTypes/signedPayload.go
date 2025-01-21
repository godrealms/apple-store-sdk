package _ResponseTypes

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"strings"
)

type signedPayload string

// DecodedPayload Decrypt structure contents
func (sp signedPayload) DecodedPayload() (*ResponseBodyV2DecodedPayload, error) {
	// Delimiter information
	header, payloadBytes, signature, err := sp.parseSignedPayload()
	if err != nil {
		return nil, err
	}

	// Get public key information
	certificate, err := header.X5c.GetPublicKey()
	if err != nil {
		return nil, err
	}

	var payload ResponseBodyV2DecodedPayload
	if err = json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse payload JSON: %v", err)
	}

	singPayload := string(sp)
	signedContent := singPayload[:strings.LastIndex(singPayload, ".")]

	// Create a hash of the signed content
	hash := sha256.Sum256([]byte(signedContent))

	// Verify the signature
	err = rsa.VerifyPKCS1v15(certificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, hash[:], signature)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %v", err)
	}

	return &payload, nil
}

// Parse signedPayload and return Header, Payload and Signature
func (sp signedPayload) parseSignedPayload() (*types.JWSDecodedHeader, []byte, []byte, error) {
	// Split signedPayload
	parts := strings.Split(string(sp), ".")
	if len(parts) != 3 {
		return nil, nil, nil, fmt.Errorf("invalid signedPayload format: expected 3 parts, got %d", len(parts))
	}

	// Decode Header
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode header: %w", err)
	}

	var header types.JWSDecodedHeader
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	// 解码 Payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// 解码 Signature
	signatureBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode signature: %w", err)
	}

	return &header, payloadBytes, signatureBytes, nil
}
