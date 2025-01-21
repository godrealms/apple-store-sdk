package TypeAliases

type RelationshipsSubscriptionGroupLocalizations struct {
	Data  []SubscriptionGroupLocalizationsData `json:"data"`
	Links RelationshipLinks                    `json:"links"`
	Meta  PagingInformation                    `json:"meta"`
}
