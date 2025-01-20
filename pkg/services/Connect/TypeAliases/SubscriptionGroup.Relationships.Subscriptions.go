package TypeAliases

// SubscriptionGroupRelationshipsSubscriptions SubscriptionGroup.Relationships.Subscriptions.Data
type SubscriptionGroupRelationshipsSubscriptions struct {
	Data  []SubscriptionGroupLocalizationsData `json:"data"`
	Links RelationshipLinks                    `json:"links"`
	Meta  PagingInformation                    `json:"meta"`
}
