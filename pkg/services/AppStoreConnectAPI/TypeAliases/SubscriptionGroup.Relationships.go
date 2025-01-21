package TypeAliases

type SubscriptionGroupRelationships struct {
	SubscriptionGroupLocalizations SubscriptionGroupLocalizations `json:"subscriptionGroupLocalizations"`
	Subscriptions                  Subscriptions                  `json:"subscriptions"`
}

type LocalizationRelationships struct {
	Data SubscriptionGroupLocalizationsData `json:"data"`
}

type SubscriptionGroupLocalizationRelationships struct {
	SubscriptionGroup LocalizationRelationships `json:"subscriptionGroup"`
}
