package errdefs

type AppleStoreAPIError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func (e *AppleStoreAPIError) Error() string {
	return e.ErrorMessage
}

func NewAppleStoreAPIError(errorCode int, errorMessage string) *AppleStoreAPIError {
	return &AppleStoreAPIError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

var (
	// AccountNotFoundError An error that indicates the App Store account wasn’t found.
	AccountNotFoundError = NewAppleStoreAPIError(4040001, "Account not found.")

	// AppNotFoundError An error that indicates the app wasn’t found.
	AppNotFoundError = NewAppleStoreAPIError(4040003, "App not found.")

	// FamilySharedSubscriptionExtensionIneligibleError An error that indicates a subscription isn’t directly eligible for a renewal date extension because the customer obtained it through Family Sharing.
	FamilySharedSubscriptionExtensionIneligibleError = NewAppleStoreAPIError(4030007, "Subscriptions that users obtain through Family Sharing can't get a renewal date extension directly.")

	// GeneralInternalError An error that indicates a general internal error.
	GeneralInternalError = NewAppleStoreAPIError(5000000, "An unknown error occurred.")

	// GeneralBadRequestError An error that indicates an invalid request.
	GeneralBadRequestError = NewAppleStoreAPIError(4000000, "Bad request.")

	// InvalidAppIdentifierError An error that indicates an invalid app identifier.
	InvalidAppIdentifierError = NewAppleStoreAPIError(4000002, "Invalid request app identifier.")

	// InvalidEmptyStorefrontCountryCodeListError An error that indicates a required storefront country code is empty.
	InvalidEmptyStorefrontCountryCodeListError = NewAppleStoreAPIError(4000027, "Invalid request. If provided, the list of storefront country codes must not be empty.")

	// InvalidExtendByDaysError An error that indicates an invalid extend-by-days value.
	InvalidExtendByDaysError = NewAppleStoreAPIError(4000009, "Invalid extend by days value.")

	// InvalidExtendReasonCodeError An error that indicates an invalid reason code.
	InvalidExtendReasonCodeError = NewAppleStoreAPIError(4000010, "Invalid extend reason code.")

	// InvalidOriginalTransactionIdError An error that indicates an invalid original transaction identifier.
	InvalidOriginalTransactionIdError = NewAppleStoreAPIError(4000008, "Invalid original transaction id.")

	// InvalidRefundPreferenceError An error that indicates an invalid refund preference code.
	InvalidRefundPreferenceError = NewAppleStoreAPIError(4000044, "Invalid request. The refund preference field is invalid.")

	// InvalidRequestIdentifierError An error that indicates an invalid request identifier.
	InvalidRequestIdentifierError = NewAppleStoreAPIError(4000011, "Invalid request identifier.")

	// InvalidRequestRevisionError An error that indicates an invalid request revision.
	InvalidRequestRevisionError = NewAppleStoreAPIError(4000005, "Invalid request revision.")

	// InvalidRevokedError An error that indicates the revoked parameter contains an invalid value.
	InvalidRevokedError = NewAppleStoreAPIError(4000030, "Invalid request. The revoked parameter is invalid.")

	// InvalidStatusError An error that indicates the status parameter is invalid.
	InvalidStatusError = NewAppleStoreAPIError(4000031, "Invalid request. The status parameter is invalid.")

	// InvalidStorefrontCountryCodeError An error that indicates a storefront code is invalid.
	InvalidStorefrontCountryCodeError = NewAppleStoreAPIError(4000028, "Invalid request. A storefront country code was invalid.")

	// InvalidTransactionIdError An error that indicates an invalid transaction identifier.
	InvalidTransactionIdError = NewAppleStoreAPIError(4000006, "Invalid transaction id.")

	// OriginalTransactionIdNotFoundError An error that indicates an original transaction identifier wasn’t found.
	OriginalTransactionIdNotFoundError = NewAppleStoreAPIError(4040005, "Original transaction id not found.")

	// RateLimitExceededError An error that indicates the request exceeded the rate limit.
	RateLimitExceededError = NewAppleStoreAPIError(4290000, "Rate limit exceeded.")

	// StatusRequestNotFoundError An error that indicates the server didn’t find a subscription-renewal-date extension
	// request for the request identifier and product identifier you provided.
	StatusRequestNotFoundError = NewAppleStoreAPIError(4040009, "The server didn't find a subscription-renewal-date extension request for this requestIdentifier and productId combination.")

	// SubscriptionExtensionIneligibleError An error that indicates the subscription doesn’t qualify for a renewal-date extension due to its subscription state.
	SubscriptionExtensionIneligibleError = NewAppleStoreAPIError(4030004, "Forbidden - subscription state ineligible for extension.")

	// SubscriptionMaxExtensionError An error that indicates the subscription doesn’t qualify for a renewal-date extension
	// because it has already received the maximum extensions.
	SubscriptionMaxExtensionError = NewAppleStoreAPIError(4030005, "Forbidden - subscription has reached maximum extension count.")

	// TransactionIdNotFoundError An error that indicates a transaction identifier wasn’t found.
	TransactionIdNotFoundError = NewAppleStoreAPIError(4040010, "Transaction id not found.")
)
