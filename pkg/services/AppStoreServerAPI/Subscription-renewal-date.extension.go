package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/RequestTypes"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"net/http"
)

type SubscriptionRenewalDateExtension struct {
	client *client.AppStoreServerAPIClient
}

func NewSubscriptionRenewalDateExtension(client *client.AppStoreServerAPIClient) *SubscriptionRenewalDateExtension {
	return &SubscriptionRenewalDateExtension{client: client}
}

// ExtendSubscriptionRenewalDate Extends the renewal date of a customer’s active subscription using the original transaction identifier.
func (rde *SubscriptionRenewalDateExtension) ExtendSubscriptionRenewalDate(originalTransactionId string, request *RequestTypes.ExtendRenewalDateRequest) (*ResponseTypes.ExtendRenewalDateResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/subscriptions/extend/%s", originalTransactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := rde.client.PUT(endpoint, headers, jsonBody)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.ExtendRenewalDateResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// ExtendSubscriptionRenewalDatesForAllActiveSubscribers Uses a subscription’s product identifier to extend the renewal date for all of its eligible active subscribers.
func (rde *SubscriptionRenewalDateExtension) ExtendSubscriptionRenewalDatesForAllActiveSubscribers(request *RequestTypes.MassExtendRenewalDateRequest) (*ResponseTypes.MassExtendRenewalDateResponse, error) {
	endpoint := "/inApps/v1/subscriptions/extend/mass/"
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := rde.client.Post(endpoint, jsonBody, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.MassExtendRenewalDateResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// GetStatusOfSubscriptionRenewalDateExtensions Checks whether a renewal date extension request completed, and provides the final count of successful or failed extensions.
func (rde *SubscriptionRenewalDateExtension) GetStatusOfSubscriptionRenewalDateExtensions(productId types.ProductId, requestIdentifier types.RequestIdentifier) (*ResponseTypes.MassExtendRenewalDateStatusResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/subscriptions/extend/mass/%s/%s", productId, requestIdentifier)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := rde.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.MassExtendRenewalDateStatusResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
