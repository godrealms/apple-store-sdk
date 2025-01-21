package AppStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
	"net/http"
)

type RefundService struct {
	client *client.ServerClient
}

func NewRefundService(client *client.ServerClient) *RefundService {
	return &RefundService{
		client: client,
	}
}

// GetRefundHistory Get a paginated list of all of a customer’s refunded in-app purchases for your app.
func (rs *RefundService) GetRefundHistory(transactionId string) (*models.RefundHistoryResponse, error) {
	endpoint := fmt.Sprintf("inApps/v2/refund/lookup/%s", transactionId)
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
		var response models.RefundHistoryResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
