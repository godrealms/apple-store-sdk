package _Dictionaries

import (
	"github.com/godrealms/apple-store-sdk/pkg/types"
)

type UnifiedReceipt struct {
	// The environment for which the App Store generated the receipt.
	Environment types.Environment `json:"environment"`
	// The latest Base64-encoded app receipt.
	LatestReceipt byte `json:"latest_receipt"`
	// An array that contains the latest 100 in-app purchase transactions of the decoded value in latest_receipt.
	// This array excludes transactions for consumable products your app has marked as finished.
	// The contents of this array are identical to those in responseBody.
	// Latest_receipt_info in the verifyReceipt endpoint response for receipt validation.
	LatestReceiptInfo []UnifiedReceiptInfoLatest `json:"latest_receipt_info"`
	// An array in which each element contains the pending renewal information for each auto-renewable subscription identified in product_id.
	// The contents of this array are identical to those in responseBody.
	// Pending_renewal_info in the verifyReceipt endpoint response for receipt validation.
	PendingRenewalInfo []UnifiedReceiptPendingRenewalInfo `json:"pending_renewal_info"`
	// The status code, where 0 indicates that the notification is valid.
	Status types.Status `json:"status"` // Value: 0
}
