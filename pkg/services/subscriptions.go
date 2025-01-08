package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"net/http"
)

// SubscriptionsService handles Apple subscription-related API calls
type SubscriptionsService struct {
	client *client.Client
}

// NewSubscriptionsService creates a new instance of SubscriptionsService
func NewSubscriptionsService(client *client.Client) *SubscriptionsService {
	return &SubscriptionsService{
		client: client,
	}
}

// VerifyReceipt verifies the receipt with Apple's App Store Server API v2
func (s *SubscriptionsService) VerifyReceipt(receipt string) (*models.ReceiptResponse, error) {
	url := fmt.Sprintf("%s/inApps/v2/verifyReceipt", s.client.BaseURL)

	// Create request payload
	payload := map[string]string{
		"receipt-data": receipt,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize receipt: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.Config.AppleJWT))

	// Execute HTTP request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var response models.ReceiptResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

// GetSubscriptionStatus retrieves the subscription status using Apple's App Store Server API v2
func (s *SubscriptionsService) GetSubscriptionStatus(originalTransactionID string) (*models.SubscriptionStatusResponse, error) {
	url := fmt.Sprintf("%s/inApps/v2/subscriptions/%s", s.client.BaseURL, originalTransactionID)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.Config.AppleJWT))

	// Execute HTTP request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var response models.SubscriptionStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

// CancelSubscription sends a request to Apple's App Store Server API v2 to cancel a subscription
func (s *SubscriptionsService) CancelSubscription(originalTransactionID string, reason int) (*models.CancelResponse, error) {
	url := fmt.Sprintf("%s/inApps/v2/subscriptions/%s/cancel", s.client.BaseURL, originalTransactionID)

	// Create request payload
	payload := map[string]int{
		"cancellation_reason": reason,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payload: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.client.Config.AppleJWT))

	// Execute HTTP request
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Parse response
	var response models.CancelResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}
