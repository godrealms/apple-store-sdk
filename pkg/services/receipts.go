package services

// ReceiptValidationRequest represents the request payload for receipt validation
type ReceiptValidationRequest struct {
	ReceiptData string `json:"receipt-data"`
}

// ReceiptValidationResponse represents the response from the receipt validation API
type ReceiptValidationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
