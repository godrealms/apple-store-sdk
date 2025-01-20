package TypeAliases

type SubscriptionGroupLocalizations string

const (
	SubscriptionGroupLocalizationsName              SubscriptionGroupLocalizations = "name"
	SubscriptionGroupLocalizationsCustomAppName     SubscriptionGroupLocalizations = "customAppName"
	SubscriptionGroupLocalizationsLocale            SubscriptionGroupLocalizations = "locale"
	SubscriptionGroupLocalizationsState             SubscriptionGroupLocalizations = "state"
	SubscriptionGroupLocalizationsSubscriptionGroup SubscriptionGroupLocalizations = "subscriptionGroup"
)

func (obj SubscriptionGroupLocalizations) String() string {
	return string(obj)
}

type SubscriptionGroupLocalization struct {
	Id            string                                     `json:"id"`
	Attributes    SubscriptionGroupLocalizationAttributes    `json:"attributes"`
	Links         ResourceLinks                              `json:"links"`
	Relationships SubscriptionGroupLocalizationRelationships `json:"relationships"`
	Type          string                                     `json:"type"` // Value: subscriptionGroupLocalizations
}

type SubscriptionGroupLocalizationsResponse struct {
	Data     []SubscriptionGroupLocalization `json:"data"`
	Included []SubscriptionGroup             `json:"included"`
	Links    PagedDocumentLinks              `json:"links"`
	Meta     PagingInformation               `json:"meta"`
}
