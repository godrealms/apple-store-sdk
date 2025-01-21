package RequestTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

// ExtendRenewalDateRequest The request body that contains subscription-renewal-extension data for an individual subscription.
type ExtendRenewalDateRequest struct {
	//Required. The number of days to extend the subscription renewal date.
	//Maximum: 90
	ExtendByDays types.ExtendByDays `json:"extendByDays"`
	//Required. The reason code for the subscription date extension.
	ExtendReasonCode types.ExtendReasonCode `json:"extendReasonCode"`
	//Required. A string that contains a value you provide to uniquely identify this renewal-date extension request.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`
}
