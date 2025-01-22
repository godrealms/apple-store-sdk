package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/QueryParameters"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"net/http"
)

type SubscriptionStatus struct {
	client *client.AppStoreServerAPIClient
}

func NewSubscriptionStatus(client *client.AppStoreServerAPIClient) *SubscriptionStatus {
	return &SubscriptionStatus{client: client}
}

// GetAllSubscriptionStatuses
// Get the statuses for all of a customer’s auto-renewable subscriptions in your app.
func (ss *SubscriptionStatus) GetAllSubscriptionStatuses(transactionId types.TransactionId, parameters QueryParameters.SubscriptionsQueryParameters) (*ResponseTypes.StatusResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/subscriptions/%s", transactionId)
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
		var response ResponseTypes.StatusResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
