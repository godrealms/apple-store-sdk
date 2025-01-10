package models

// ConsumptionRequest The request body containing consumption information.
type ConsumptionRequest struct {
	// (Required) The age of the customer’s account.
	AccountTenure accountTenure `json:"accountTenure"`

	// (Required) The UUID of the in-app user account that completed the in-app purchase transaction.
	AppAccountToken appAccountToken `json:"appAccountToken"`

	// (Required) A value that indicates the extent to which the customer consumed the in-app purchase.
	ConsumptionStatus consumptionStatus `json:"consumptionStatus"`

	// (Required) A Boolean value of true or false that indicates whether the customer consented to provide consumption data.
	//Note: The App Store server rejects requests that have a customerConsented value other than true by returning an HTTP 400 error with an InvalidCustomerConsentError.
	CustomerConsented customerConsented `json:"customerConsented"`

	// (Required) A value that indicates whether the app successfully delivered an in-app purchase that works properly.
	DeliveryStatus deliveryStatus `json:"deliveryStatus"`

	// (Required) A value that indicates the total amount, in USD, of in-app purchases the customer has made in your app, across all platforms.
	LifetimeDollarsPurchased lifetimeDollarsPurchased `json:"lifetimeDollarsPurchased"`

	// (Required) A value that indicates the total amount, in USD, of refunds the customer has received, in your app, across all platforms.
	LifetimeDollarsRefunded lifetimeDollarsRefunded `json:"lifetimeDollarsRefunded"`

	// (Required) A value that indicates the platform on which the customer consumed the in-app purchase.
	Platform platform `json:"platform"`

	// (Required) A value that indicates the amount of time that the customer used the app.
	PlayTime playTime `json:"playTime"`

	// A value that indicates your preference, based on your operational logic, as to whether Apple should grant the refund.
	RefundPreference refundPreference `json:"refundPreference"`

	// (Required) A Boolean value of true or false that indicates whether you provided, prior to its purchase,
	//a free sample or trial of the content, or information about its functionality.
	SampleContentProvided sampleContentProvided `json:"sampleContentProvided"`

	// (Required) The status of the customer’s account.
	UserStatus userStatus `json:"userStatus"`
}
