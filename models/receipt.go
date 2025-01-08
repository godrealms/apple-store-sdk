package models

// ReceiptResponse represents the response from Apple's receipt verification API
type ReceiptResponse struct {
	Status            int    `json:"status"`
	Environment       string `json:"environment"`
	LatestReceipt     string `json:"latest_receipt"`
	LatestReceiptInfo []struct {
		OriginalTransactionID string `json:"original_transaction_id"`
		ProductID             string `json:"product_id"`
		PurchaseDate          string `json:"purchase_date"`
		ExpirationDate        string `json:"expires_date"`
	} `json:"latest_receipt_info"`
}
