package services

import (
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
)

type RefundService struct {
	client *client.Client
}

func NewRefundService(client *client.Client) *RefundService {
	return &RefundService{
		client: client,
	}
}

// GetRefundHistory Get a paginated list of all of a customer’s refunded in-app purchases for your app.
func (rs *RefundService) GetRefundHistory(transactionId string) (*models.RefundHistoryResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v2/refund/lookup/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := rs.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	_ = body
	_ = code
	return nil, nil
}
