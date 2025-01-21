package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

// MassExtendRenewalDateResponse A response that indicates the server successfully received the subscription-renewal-date extension request.
type MassExtendRenewalDateResponse struct {
	// A string that contains the UUID that identifies the subscription-renewal-date extension request.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`
}
