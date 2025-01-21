package QueryParameters

import "github.com/godrealms/apple-store-sdk/pkg/types"

type SubscriptionsQueryParameters struct {
	// An optional filter that indicates the status of subscriptions to include in the response.
	// Your query may specify more than one status query parameter.
	Status types.Status `json:"status"`
}
