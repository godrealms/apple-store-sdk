package services

import (
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
)

type TransactionsService struct {
	client *client.Client
}

func NewTransactionsService(client *client.Client) *TransactionsService {
	return &TransactionsService{
		client: client,
	}
}

func (ts *TransactionsService) GetHistory(transactionId string) (*models.TransactionHistory, error) {
	url := fmt.Sprintf("/inApps/v2/history/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ts.client.Get(url, headers, nil)
	if err != nil {
		return nil, err
	}
	_ = body
	_ = code
	return nil, nil
}

func (ts *TransactionsService) GetInfo(transactionId string) (*models.TransactionInfo, error) {
	url := fmt.Sprintf("/inApps/v1/transactions/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ts.client.Get(url, headers, nil)
	if err != nil {
		return nil, err
	}
	_ = body
	_ = code
	return nil, nil
}
