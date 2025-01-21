package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type OrderLookupResponse struct {
	// The status that indicates whether the order ID is valid.
	Status types.OrderLookupStatus `json:"status"`
	// An array of in-app purchase transactions that are part of order, signed by Apple, in JSON Web Signature format.
	SignedTransactions []types.JWSTransaction `json:"signedTransactions"`
}
