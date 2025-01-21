package AppStoreServerAPI

import "github.com/godrealms/apple-store-sdk/pkg/client"

type SubscriptionStatus struct {
	client *client.ServerClient
}

func NewSubscriptionStatus(client *client.ServerClient) *SubscriptionStatus {
	return &SubscriptionStatus{client: client}
}

// GetAllSubscriptionStatuses Get the statuses for all of a customer’s auto-renewable subscriptions in your app.
func (s *SubscriptionStatus) GetAllSubscriptionStatuses() {}
