package errdefs

var (
	// InvalidAccountTenureError An error that indicates the value of the account tenure field is invalid.
	InvalidAccountTenureError = NewAppleStoreAPIError(4000032, "Invalid request. The account tenure field is invalid.")

	// InvalidAppAccountTokenError An error that indicates the value of the app account token field is invalid.
	InvalidAppAccountTokenError = NewAppleStoreAPIError(4000033, "Invalid request. The app account token field must contain a valid UUID or an empty string.")

	// InvalidConsumptionStatusError An error that indicates the value of the consumption status field is invalid.
	InvalidConsumptionStatusError = NewAppleStoreAPIError(4000034, "Invalid request. The consumption status field is invalid.")

	// InvalidCustomerConsentedError An error that indicates the customer consented field is invalid or doesn’t indicate that the customer consented.
	InvalidCustomerConsentedError = NewAppleStoreAPIError(4000035, "Invalid request. The customer consented field is required and must indicate the customer consented.")

	// InvalidDeliveryStatusError An error that indicates the value in the delivery status field is invalid.
	InvalidDeliveryStatusError = NewAppleStoreAPIError(4000036, "Invalid request. The delivery status field is invalid.")

	// InvalidLifetimeDollarsPurchasedError An error that indicates the value in the lifetime dollars purchased field is invalid.
	InvalidLifetimeDollarsPurchasedError = NewAppleStoreAPIError(4000037, "Invalid request. The lifetime dollars purchased field is invalid.")

	// InvalidLifetimeDollarsRefundedError An error that indicates the value in the lifetime dollars refunded field is invalid.
	InvalidLifetimeDollarsRefundedError = NewAppleStoreAPIError(4000038, "Invalid request. The lifetime dollars refunded field is invalid.")

	// InvalidPlatformError An error that indicates the value in the platform field is invalid.
	InvalidPlatformError = NewAppleStoreAPIError(4000039, "Invalid request. The platform field is invalid.")
)
