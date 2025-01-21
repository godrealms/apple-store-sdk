package errdefs

var (
	// InvalidEndDateError An error that indicates the end date is invalid.
	InvalidEndDateError = NewAppleStoreAPIError(4000016, "Invalid request. The end date is not a timestamp value represented in milliseconds.")

	// InvalidNotificationTypeError An error that indicates the notification type or subtype is invalid.
	// For a list of valid notification types and subtypes, see notificationType and subtype. If you provide both a notification type and subtype, ensure the two are a valid combination, as documented on subtype.
	InvalidNotificationTypeError = NewAppleStoreAPIError(4000018, "Invalid request. The notification type or subtype is invalid.")

	// InvalidPaginationTokenError An error that indicates the pagination token is invalid.
	InvalidPaginationTokenError = NewAppleStoreAPIError(4000014, "Invalid request. The pagination token is invalid.")

	// InvalidStartDateError An error that indicates the start date is invalid.
	InvalidStartDateError = NewAppleStoreAPIError(4000015, "Invalid request. The start date is not a timestamp value represented in milliseconds.")

	// InvalidTestNotificationTokenError An error that indicates the test notification token is invalid.
	InvalidTestNotificationTokenError = NewAppleStoreAPIError(4000020, "Invalid request. The test notification token is invalid.")

	// InvalidInAppOwnershipTypeError An error that indicates an invalid in-app ownership type parameter.
	InvalidInAppOwnershipTypeError = NewAppleStoreAPIError(4000026, "Invalid request. The in-app ownership type parameter is invalid.")

	// InvalidProductIdError An error that indicates the product ID parameter is invalid.
	InvalidProductIdError = NewAppleStoreAPIError(4000023, "Invalid request. The product id parameter is invalid.")

	// InvalidProductTypeError An error that indicates the product type parameter is invalid.
	InvalidProductTypeError = NewAppleStoreAPIError(4000022, "An error that indicates the product type parameter is invalid.")

	// InvalidSortError An error that indicates the sort parameter is invalid.
	InvalidSortError = NewAppleStoreAPIError(4000021, "Invalid request. The sort parameter is invalid.")

	// InvalidSubscriptionGroupIdentifierError An error that indicates the subscription group identifier is invalid.
	InvalidSubscriptionGroupIdentifierError = NewAppleStoreAPIError(4000024, "Invalid request. The subscription group identifier parameter is invalid.")

	// MultipleFiltersSuppliedError An error that indicates the request is invalid because it has too many applied constraints.
	MultipleFiltersSuppliedError = NewAppleStoreAPIError(4000019, "Invalid request. Supply either a transaction id or a notification type, but not both.")

	// PaginationTokenExpiredError An error that indicates the pagination token expired.
	PaginationTokenExpiredError = NewAppleStoreAPIError(4000017, "Invalid request. The pagination token is expired.")

	// ServerNotificationURLNotFoundError An error that indicates the App Store server couldn’t find a notifications URL for your app in the environment.
	ServerNotificationURLNotFoundError = NewAppleStoreAPIError(4040007, "No App Store AppStoreServerAPI Notification URL found for provided app. Check that a URL is configured in App Store AppStoreConnectAPI for this environment.")

	// StartDateAfterEndDateError An error that indicates the end date precedes the start date, or the two dates are equal.
	StartDateAfterEndDateError = NewAppleStoreAPIError(4000013, "Invalid request. The end date precedes the start date or the dates are the same.")

	// StartDateTooFarInPastError An error that indicates the start date is earlier than the earliest allowed date.
	StartDateTooFarInPastError = NewAppleStoreAPIError(4000012, "Invalid request. The start date is earlier than the allowed start date.")

	// TestNotificationNotFoundError An error that indicates the test notification token is expired or the test notification status isn’t available.
	TestNotificationNotFoundError = NewAppleStoreAPIError(4040008, "Either the test notification token is expired or the notification and status are not yet available.")
)
