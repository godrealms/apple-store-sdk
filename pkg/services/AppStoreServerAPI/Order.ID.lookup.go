package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"net/http"
)

type LookUpOrderID struct {
	client *client.AppStoreServerAPIClient
}

func NewLookUpOrderID(client *client.AppStoreServerAPIClient) *LookUpOrderID {
	return &LookUpOrderID{
		client: client,
	}
}

// LookUpOrderID Get a customer’s in-app purchases from a receipt using the order ID.
func (os *LookUpOrderID) LookUpOrderID(orderId string) (*ResponseTypes.OrderLookupResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/lookup/%s", orderId)
	body, code, err := os.client.Get(endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.OrderLookupResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
