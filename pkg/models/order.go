package models

type OrderLookupResponse struct {
	// The status that indicates whether the order ID is valid.
	Status OrderLookupStatus `json:"status"`
	// An array of in-app purchase transactions that are part of order, signed by Apple, in JSON Web Signature format.
	SignedTransactions []JWSTransaction `json:"signedTransactions"`
}
