package AppStore

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
)

type Subscriptions struct {
	client *client.ConnectClient
}

func NewSubscriptions(client *client.ConnectClient) *SubscriptionGroups {
	client.Config.BaseURL = "https://api.appstoreconnect.apple.com"
	return &SubscriptionGroups{client: client}
}

// CreateAutoRenewableSubscription Create an auto-renewable subscription for your app.
func (s *Subscriptions) CreateAutoRenewableSubscription() {}

// ReadSubscriptionInformation Get information about a specific auto-renewable subscription.
func (s *Subscriptions) ReadSubscriptionInformation() {}

// ModifyAutoRenewableSubscription Update a specific auto-renewable subscription.
func (s *Subscriptions) ModifyAutoRenewableSubscription() {}

// DeleteSubscription Delete a specific auto-renewable subscription that you configured for an app.
func (s *Subscriptions) DeleteSubscription() {}

// ListAllLocalizationsForAutoRenewableSubscription Get a list of the subscription localizations for a specific auto-renewable subscription.
func (s *Subscriptions) ListAllLocalizationsForAutoRenewableSubscription() {}

// ListAllIntroductoryOffersForSubscription Get a list of introductory offers for a specific auto-renewable subscription.
func (s *Subscriptions) ListAllIntroductoryOffersForSubscription() {}

// ListAllIntroductoryOfferResourceIDsForAutoRenewableSubscription Get a list of resource IDs representing introductory offers for an auto-renewable subscription.
func (s *Subscriptions) ListAllIntroductoryOfferResourceIDsForAutoRenewableSubscription() {}

// DeleteIntroductoryOfferFromSubscription Delete a specific introductory offer for an auto-renewable subscription.
func (s *Subscriptions) DeleteIntroductoryOfferFromSubscription() {}

// ReadPromotedPurchaseInformationForSubscription Get details about the promoted purchase of an auto-renewable subscription.
func (s *Subscriptions) ReadPromotedPurchaseInformationForSubscription() {}

// ListAllOfferCodesForSubscription Get a list of subscription offer codes for a specific auto-renewable subscription.
func (s *Subscriptions) ListAllOfferCodesForSubscription() {}

// ListAllPromotionalOfferResourceIDsForAutoRenewableSubscription Get a list of promotional offers for a specific auto-renewable subscription.
func (s *Subscriptions) ListAllPromotionalOfferResourceIDsForAutoRenewableSubscription() {}

// ListAllPricePointsForSubscription Get a list of price points for an auto-renewable subscription by territory.
func (s *Subscriptions) ListAllPricePointsForSubscription() {}

// ListAllPricesForSubscription Get a list of prices for an auto-renewable subscription, by territory.
func (s *Subscriptions) ListAllPricesForSubscription() {}

// ListAllSubscriptionPricesIDsForAutoRenewableSubscription Get a list of resource IDs representing subscription prices for an auto-renewable subscription.
func (s *Subscriptions) ListAllSubscriptionPricesIDsForAutoRenewableSubscription() {}

// DeletePricesFromSubscription Delete a scheduled subscription price change for an auto-renewable subscription.
func (s *Subscriptions) DeletePricesFromSubscription() {}

// ReadReviewScreenshotInformationForSubscription Get information about review screenshot for a specific auto-renewable subscription.
func (s *Subscriptions) ReadReviewScreenshotInformationForSubscription() {}

// ReadInformationAboutTheAvailabilityOfSubscription Get information about the territory availability for a subscription.
func (s *Subscriptions) ReadInformationAboutTheAvailabilityOfSubscription() {}
