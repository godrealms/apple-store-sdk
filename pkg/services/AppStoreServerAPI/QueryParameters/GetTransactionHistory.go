package QueryParameters

import "github.com/godrealms/apple-store-sdk/pkg/types"

type TransactionHistoryQueryParameters struct {
	// A token you provide to get the next set of up to 20 transactions.
	// All responses include a revision token. Use the revision token from the previous HistoryResponse.
	// Note: The revision token is required in all requests except the initial request.
	// For requests that use the revision token, include the same query parameters from the initial request.
	Revision types.Revision `json:"revision"`
	// An optional start date of the timespan for the transaction history records you’re requesting.
	// The startDate needs to precede the endDate if you specify both dates.
	// The results include a transaction if its purchaseDate is equal to or greater than the startDate.
	StartDate types.Timestamp `json:"startDate"`
	// An optional end date of the timespan for the transaction history records you’re requesting.
	// Choose an endDate that’s later than the startDate if you specify both dates.
	// Using an endDate in the future is valid.
	// The results include a transaction if its purchaseDate is less than the endDate.
	EndDate types.Timestamp `json:"endDate"`
	// An optional filter that indicates the product identifier to include in the transaction history.
	// Your query may specify more than one productID.
	ProductId []types.ProductId `json:"productId"`
	// An optional filter that indicates the product type to include in the transaction history.
	// Your query may specify more than one productType.
	// Possible Values: AUTO_RENEWABLE, NON_RENEWABLE, CONSUMABLE, NON_CONSUMABLE
	ProductType []types.ProductType `json:"productType"`
	// An optional filter that limits the transaction history by the in-app ownership type.
	InAppOwnershipType types.InAppOwnershipType `json:"inAppOwnershipType"`
	// An optional sort order for the transaction history records.
	// The response sorts the transaction records by their recently modified date.
	// The default value is ASCENDING, so you receive the oldest records first.
	// Possible Values: ASCENDING, DESCENDING
	Sort string `json:"sort"`
	// An optional Boolean value that indicates whether the response includes only revoked transactions when the value is true,
	// or contains only nonrevoked transactions when the value is false.
	// By default, the request doesn’t include this parameter.
	// Possible Values: true, false
	Revoked bool `json:"revoked"`
	// An optional filter that indicates the subscription group identifier to include in the transaction history.
	// Your query may specify more than one subscriptionGroupIdentifier.
	SubscriptionGroupIdentifier types.SubscriptionGroupIdentifier `json:"subscriptionGroupIdentifier"`
	// Deprecated  Set revoked to false to exclude revoked transactions instead.
	// Possible Values: true, false
	ExcludeRevoked bool `json:"excludeRevoked,omitempty"`
}
