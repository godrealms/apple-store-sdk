package ResponseTypes

import "github.com/godrealms/apple-store-sdk/pkg/types"

// MassExtendRenewalDateStatusResponse A response that indicates the current status of a request to extend the subscription renewal date to all eligible subscribers.
type MassExtendRenewalDateStatusResponse struct {
	//The UUID that represents your request for a subscription-renewal-date extension.
	//Maximum length: 128
	RequestIdentifier types.RequestIdentifier `json:"requestIdentifier"`

	//A Boolean value that’s TRUE to indicate that the App Store completed your request to extend a subscription renewal date for all eligible subscribers.
	//The value is FALSE if the request is in progress.
	Complete types.Complete `json:"complete"`

	//The date that the App Store completes the request.
	//Appears only if complete is TRUE.
	CompleteDate types.Timestamp `json:"completeDate"`

	//The final count of subscribers that fail to receive a subscription-renewal-date extension.
	//Appears only if complete is TRUE.
	FailedCount types.FailedCount `json:"failedCount"`

	//The final count of subscribers that successfully receive a subscription-renewal-date extension.
	//Appears only if complete is TRUE.
	SucceededCount types.SucceededCount `json:"succeededCount"`
}
