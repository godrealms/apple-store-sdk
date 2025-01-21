package TypeAliases

type Subscriptions string

const (
	SubscriptionsName                      Subscriptions = "name"
	SubscriptionsProductId                 Subscriptions = "productId"
	SubscriptionsFamilySharable            Subscriptions = "familySharable"
	SubscriptionsState                     Subscriptions = "state"
	SubscriptionsSubscriptionPeriod        Subscriptions = "subscriptionPeriod"
	SubscriptionsReviewNote                Subscriptions = "reviewNote"
	SubscriptionsGroupLevel                Subscriptions = "groupLevel"
	SubscriptionsSubscriptionLocalizations Subscriptions = "subscriptionLocalizations"
	SubscriptionsAppStoreReviewScreenshot  Subscriptions = "appStoreReviewScreenshot"
	SubscriptionsGroup                     Subscriptions = "group"
	SubscriptionsIntroductoryOffers        Subscriptions = "introductoryOffers"
	SubscriptionsPromotionalOffers         Subscriptions = "promotionalOffers"
	SubscriptionsOfferCodes                Subscriptions = "offerCodes"
	SubscriptionsPrices                    Subscriptions = "prices"
	SubscriptionsPricePoints               Subscriptions = "pricePoints"
	SubscriptionsPromotedPurchase          Subscriptions = "promotedPurchase"
	SubscriptionsSubscriptionAvailability  Subscriptions = "subscriptionAvailability"
	SubscriptionsWinBackOffers             Subscriptions = "winBackOffers"
	SubscriptionsImages                    Subscriptions = "images"
)

func (s Subscriptions) String() string {
	return string(s)
}
