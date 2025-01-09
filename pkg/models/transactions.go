package models

type TransactionHistory struct {
	// The app’s identifier in the App Store.
	AppAppleId appAppleId `json:"appAppleId"`

	// The bundle identifier of the app.
	BundleId bundleId `json:"bundleId"`

	// The server environment in which you’re making the request, whether sandbox or production.
	Environment environment `json:"environment"`

	// A Boolean value that indicates whether the App Store has more transactions than it returns in this response.
	//If the value is true, use the revision token to request the next set of transactions.
	HasMore hasMore `json:"hasMore"`

	// A token you use in a query to request the next set of transactions from the endpoint.
	Revision revision `json:"revision"`

	// An array of in-app purchase transactions for the customer, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactions []JWSTransaction `json:"signedTransactions"`
}

type TransactionInfo struct {
	// A customer’s in-app purchase transaction, signed by Apple, in JSON Web Signature (JWS) format.
	SignedTransactionInfo JWSTransaction `json:"signedTransactionInfo"`
}
