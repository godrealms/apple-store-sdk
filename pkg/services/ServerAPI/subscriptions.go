package ServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
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

func (ss *SubscriptionsService) GetAllSubscriptionStatuses(transactionId string, status ...string) (*models.StatusResponse, error) {
	parameters := make(map[string][]string)
	if len(status) > 0 {
		parameters["status"] = status
	}

	endpoint := fmt.Sprintf("inApps/v1/subscriptions/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ss.client.Get(endpoint, headers, parameters)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.StatusResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// ExtendSubscriptionRenewalDate Extends the renewal date of a customer’s active subscription using the original transaction identifier.
func (ss *SubscriptionsService) ExtendSubscriptionRenewalDate(originalTransactionId string, request *models.ExtendRenewalDateRequest) (*models.ExtendRenewalDateResponse, error) {
	endpoint := fmt.Sprintf("inApps/v1/subscriptions/extend/%s", originalTransactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := ss.client.PUT(endpoint, headers, jsonBody)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.ExtendRenewalDateResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// ExtendSubscriptionRenewalDatesForAllActiveSubscribers Uses a subscription’s product identifier to extend the renewal date for all of its eligible active subscribers.
func (ss *SubscriptionsService) ExtendSubscriptionRenewalDatesForAllActiveSubscribers(request *models.MassExtendRenewalDateRequest) (*models.MassExtendRenewalDateResponse, error) {
	endpoint := "inApps/v1/subscriptions/extend/mass/"
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := ss.client.Post(endpoint, jsonBody, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.MassExtendRenewalDateResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// GetStatusOfSubscriptionRenewalDateExtensions Checks whether a renewal date extension request completed, and provides the final count of successful or failed extensions.
func (ss *SubscriptionsService) GetStatusOfSubscriptionRenewalDateExtensions(productId, requestIdentifier string) (*models.MassExtendRenewalDateStatusResponse, error) {
	endpoint := fmt.Sprintf("inApps/v1/subscriptions/extend/mass/%s/%s", productId, requestIdentifier)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ss.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.MassExtendRenewalDateStatusResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
