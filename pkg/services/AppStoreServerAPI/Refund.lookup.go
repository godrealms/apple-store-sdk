package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"net/http"
)

type RefundLookup struct {
	client *client.AppStoreServerAPIClient
}

func NewRefundLookup(client *client.AppStoreServerAPIClient) *RefundLookup {
	return &RefundLookup{
		client: client,
	}
}

// GetRefundHistory Get a paginated list of all of a customer’s refunded in-app purchases for your app.
func (rs *RefundLookup) GetRefundHistory(transactionId string) (*ResponseTypes.RefundHistoryResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v2/refund/lookup/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := rs.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}

	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.RefundHistoryResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
