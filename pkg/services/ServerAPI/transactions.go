package ServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/ServerAPI/models"
	"net/http"
)

type TransactionsService struct {
	client *client.Client
}

func NewTransactionsService(client *client.Client) *TransactionsService {
	return &TransactionsService{
		client: client,
	}
}

// GetTransactionHistory Get a customer’s in-app purchase transaction history for your app.
func (ts *TransactionsService) GetTransactionHistory(transactionId string) (*models.HistoryResponse, error) {
	url := fmt.Sprintf("/inApps/v2/history/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ts.client.Get(url, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.HistoryResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// GetTransactionInfo Get information about a single transaction for your app.
func (ts *TransactionsService) GetTransactionInfo(transactionId string) (*models.TransactionInfoResponse, error) {
	url := fmt.Sprintf("/inApps/v1/transactions/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ts.client.Get(url, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.TransactionInfoResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
