package models

// SubscriptionStatusResponse represents the subscription status response
type SubscriptionStatusResponse struct {
	Data []struct {
		OriginalTransactionID string `json:"original_transaction_id"`
		ProductID             string `json:"product_id"`
		Status                int    `json:"status"`
		ExpirationDate        string `json:"expires_date"`
	} `json:"data"`
}

// CancelResponse represents the response from Apple's cancel subscription API
type CancelResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
