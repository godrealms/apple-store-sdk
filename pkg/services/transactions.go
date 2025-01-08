package services

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
)

type TransactionsService struct {
	client *client.Client
}

func NewTransactionsService(client *client.Client) *TransactionsService {
	return &TransactionsService{
		client: client,
	}
}
