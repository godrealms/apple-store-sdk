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

type InAppPurchaseHistory struct {
	client *client.AppStoreServerAPIClient
}

func NewInAppPurchaseHistory(client *client.AppStoreServerAPIClient) *InAppPurchaseHistory {
	return &InAppPurchaseHistory{client: client}
}

// GetTransactionHistory Get a customer’s in-app purchase transaction history for your app.
func (ph *InAppPurchaseHistory) GetTransactionHistory(transactionId types.TransactionId, parameters *QueryParameters.TransactionHistoryQueryParameters) (*ResponseTypes.HistoryResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v2/history/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ph.client.Get(endpoint, headers, parameters)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.HistoryResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
