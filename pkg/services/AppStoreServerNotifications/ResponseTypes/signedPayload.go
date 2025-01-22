package responseTypes

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"math/big"
	"strings"
)

type SignedPayload string

// DecodedPayload Decrypt structure contents
func (sp SignedPayload) DecodedPayload() (*ResponseBodyV2DecodedPayload, error) {
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
	switch pubKey := certificate.PublicKey.(type) {
	case *ecdsa.PublicKey: // 使用 ECDSA 验证签名
		var r, s big.Int
		r.SetBytes(signature[:len(signature)/2])
		s.SetBytes(signature[len(signature)/2:])
		if ecdsa.Verify(pubKey, hash[:], &r, &s) {
			return &payload, nil
		} else if ecdsa.VerifyASN1(pubKey, hash[:], signature) {
			return &payload, nil
		}
		return nil, fmt.Errorf("signatureBytes verification failed")
	case *rsa.PublicKey: // 使用 RSA 验证签名
		if err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signature); err != nil {
			return nil, fmt.Errorf("signature verification failed: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pubKey)
	}

	return &payload, nil
}

// Parse SignedPayload and return Header, Payload and Signature
func (sp SignedPayload) parseSignedPayload() (*types.JWSDecodedHeader, []byte, []byte, error) {
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
