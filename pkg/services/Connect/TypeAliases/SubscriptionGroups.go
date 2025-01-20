package TypeAliases

import "strconv"

type SubscriptionGroups string

const (
	SubscriptionGroupsReferenceName                  SubscriptionGroups = "referenceName"
	SubscriptionGroupsSubscriptions                  SubscriptionGroups = "subscriptions"
	SubscriptionGroupsSubscriptionGroupLocalizations SubscriptionGroups = "subscriptionGroupLocalizations"
)

func (s SubscriptionGroups) String() string {
	return string(s)
}

func (s SubscriptionGroups) Integer() int64 {
	parseInt, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return 0
	}
	return parseInt
}
