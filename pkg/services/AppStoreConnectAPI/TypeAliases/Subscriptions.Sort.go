package TypeAliases

type SubscriptionsSort string

func (s SubscriptionsSort) ASC() string {
	return "referenceName"
}

func (s SubscriptionsSort) DESC() string {
	return "-referenceName"
}
