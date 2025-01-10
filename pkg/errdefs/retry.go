package errdefs

var (
	// AccountNotFoundRetryableError An error response that indicates the App Store account wasn’t found, but you can try again.
	AccountNotFoundRetryableError = NewAppleStoreAPIError(4040002, "Account not found. Please try again.")

	// AppNotFoundRetryableError An error response that indicates the app wasn’t found, but you can try again.
	AppNotFoundRetryableError = NewAppleStoreAPIError(4040004, "App not found. Please try again.")

	// GeneralInternalRetryableError An error response that indicates an unknown error occurred, but you can try again.
	GeneralInternalRetryableError = NewAppleStoreAPIError(5000001, "An unknown error occurred. Please try again.")

	// OriginalTransactionIdNotFoundRetryableError An error response that indicates the original transaction identifier wasn’t found, but you can try again.
	OriginalTransactionIdNotFoundRetryableError = NewAppleStoreAPIError(4040006, "Original transaction id not found. Please try again.")
)
