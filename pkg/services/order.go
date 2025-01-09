package services

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
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
func (o *OrderService) LookUpOrderID(orderID string) {

}
