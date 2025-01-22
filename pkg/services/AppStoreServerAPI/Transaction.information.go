package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"net/http"
)

type TransactionInformation struct {
	client *client.AppStoreServerAPIClient
}

func NewTransactionInformation(client *client.AppStoreServerAPIClient) *TransactionInformation {
	return &TransactionInformation{client: client}
}

func (ti *TransactionInformation) GetTransactionInfo(transactionId types.TransactionId) (*ResponseTypes.TransactionInfoResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/transactions/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ti.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.TransactionInfoResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
