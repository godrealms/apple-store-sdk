package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

type HistoryResponse struct {
	// The app’s identifier in the App Store.
	AppAppleId types.AppAppleId `json:"appAppleId"`
	// The bundle identifier of the app.
	BundleId types.BundleId `json:"bundleId"`
	// The server environment in which you’re making the request, whether sandbox or production.
	Environment types.Environment `json:"environment"`
	// A Boolean value that indicates whether the App Store has more transactions than it returns in this response.
	// If the value is true, use the revision token to request the next set of transactions.
	HasMore types.HasMore `json:"hasMore"`
	// A token you use in a query to request the next set of transactions from the endpoint.
	Revision types.Revision `json:"revision"`
	// An array of in-app purchase transactions for the customer, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactions []types.JWSTransaction `json:"signedTransactions"`
}
