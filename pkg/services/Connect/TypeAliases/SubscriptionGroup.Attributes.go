package TypeAliases

type SubscriptionGroupAttributes struct {
	ReferenceName string `json:"referenceName"`
}

type SubscriptionGroupLocalizationAttributes struct {
	CustomAppName string `json:"customAppName"`
	Locale        string `json:"locale"` // The specified locale. To learn more, see Managing metadata in your app by using locale shortcodes.
	Name          string `json:"name"`
	State         string `json:"state"` // Possible Values: PREPARE_FOR_SUBMISSION, WAITING_FOR_REVIEW, APPROVED, REJECTED
}
