package TypeAliases

type SubscriptionGroupsQueryParameters struct {
	Fields  []string `json:"fields"`  // subscriptionGroupLocalizations/subscriptionGroups/subscriptions/referenceName/subscriptions.state
	Include []string `json:"include"` // Possible Values: subscriptions, subscriptionGroupLocalizations
	Limit   int64    `json:"limit"`   // Maximum: 200 / subscriptionGroupLocalizations Maximum: 50 / subscriptions Maximum: 50
	Sort    []string `json:"sort"`    // Possible Values: referenceName, -referenceName
}

type SubscriptionGroup struct {
	Id            string                         `json:"id"`
	Attributes    SubscriptionGroupAttributes    `json:"attributes"`
	Links         ResourceLinks                  `json:"links"`
	Relationships SubscriptionGroupRelationships `json:"relationships"`
}

type SubscriptionGroupsResponse struct {
	Data     []SubscriptionGroup `json:"data"`
	Included []any               `json:"included,omitempty"` // Subscription, SubscriptionGroupLocalization
	Links    PagedDocumentLinks  `json:"links"`
	Meta     PagingInformation   `json:"meta"`
}
