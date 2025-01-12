package ServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/ServerAPI/models"
	"net/http"
)

type OrderService struct {
	client *client.Client
}

func NewOrderService(client *client.Client) *OrderService {
	return &OrderService{
		client: client,
	}
}

// LookUpOrderID Get a customer’s in-app purchases from a receipt using the order ID.
func (os *OrderService) LookUpOrderID(orderId string) (*models.OrderLookupResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/lookup/%s", orderId)
	body, code, err := os.client.Get(endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.OrderLookupResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
