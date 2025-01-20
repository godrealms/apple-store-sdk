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
