package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type TransactionInfoResponse struct {
	// A customer’s in-app purchase transaction, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactionInfo types.JWSTransaction `json:"signedTransactionInfo"`
}
