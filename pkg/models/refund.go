package models

import "github.com/godrealms/apple-store-sdk/pkg/types"

// RefundHistoryResponse A response that contains an array of signed JSON Web Signature (JWS) refunded transactions, and paging information.
type RefundHistoryResponse struct {
	// A Boolean value that indicates whether the App Store has more transactions than it returns in signedTransactions.
	// If the value is true, use the revision token to request the next set of transactions by calling Get Refund History.
	HasMore types.HasMore `json:"hasMore"`

	// A token you provide in a query to request the next set of transactions from the Get Refund History endpoint.
	Revision types.Revision `json:"revision"`

	// A list of up to 20 JWS transactions, or an empty array if the customer hasn’t received any refunds in your app.
	// The transactions are sorted in ascending order by revocationDate.
	SignedTransactions []types.JWSTransaction `json:"signedTransactions"`
}
