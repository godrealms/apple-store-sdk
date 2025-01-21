package AppStore

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
)

type CreatingManagingSubscriptionOffers struct {
	client *client.ConnectClient
}

func NewCreatingManagingSubscriptionOffers(client *client.ConnectClient) *SubscriptionGroups {
	client.Config.BaseURL = "https://api.appstoreconnect.apple.com"
	return &SubscriptionGroups{client: client}
}

// CreateSubscriptionOffer Create a subscription offer that provides offer codes for an auto-renewable subscription.
func (cms *CreatingManagingSubscriptionOffers) CreateSubscriptionOffer() {}

// ReadSubscriptionOfferCodeInformation Get details about a specific subscription offer that has offer codes for an auto-renewable subscription.
func (cms *CreatingManagingSubscriptionOffers) ReadSubscriptionOfferCodeInformation() {}

// DeactivateSubscriptionOfferWithOfferCodes Deactivate a subscription offer that has offer codes for an auto-renewable subscription.
func (cms *CreatingManagingSubscriptionOffers) DeactivateSubscriptionOfferWithOfferCodes() {}

// ListAllSubscriptionOfferCodePrices Get a list of price tiers for a subscription offer code.
func (cms *CreatingManagingSubscriptionOffers) ListAllSubscriptionOfferCodePrices() {}

// SubscriptionIntroductoryOffers Create, modify, and delete introductory offers for auto-renewable subscriptions.
type SubscriptionIntroductoryOffers struct {
	client *client.ConnectClient
}

func NewSubscriptionIntroductoryOffers(client *client.ConnectClient) *SubscriptionGroups {
	return &SubscriptionGroups{client: client}
}

// CreateIntroductoryOffer Create an introductory offer for an auto-renewable subscription.
func (cms *SubscriptionIntroductoryOffers) CreateIntroductoryOffer() {}

// ModifyIntroductoryOffer Update a specific introductory offer for an auto-renewable subscription.
func (cms *SubscriptionIntroductoryOffers) ModifyIntroductoryOffer() {}

// DeleteIntroductoryOfferSubscription Delete a specific introductory offer for an auto-renewable subscription.
func (cms *SubscriptionIntroductoryOffers) DeleteIntroductoryOfferSubscription() {}

type SubscriptionPromotionalOffers struct {
	client *client.ConnectClient
}

func NewSubscriptionPromotionalOffers(client *client.ConnectClient) *SubscriptionGroups {
	return &SubscriptionGroups{client: client}
}

// CreatePromotionalOffer Create a promotional offer for an auto-renewable subscription.
func (cms *SubscriptionPromotionalOffers) CreatePromotionalOffer() {}

// ListAllPromotionalOfferPricesForSubscription Get a list of prices of a promotional offer for an auto-renewable subscription, for a specified territory.
func (cms *SubscriptionPromotionalOffers) ListAllPromotionalOfferPricesForSubscription() {}

// ReadPromotionalOfferInformation Get details about a specific promotional offer for an auto-renewable subscription.
func (cms *SubscriptionPromotionalOffers) ReadPromotionalOfferInformation() {}

// ModifyPromotionalOffer Update the prices for a specific promotional offer for an auto-renewable subscription.
func (cms *SubscriptionPromotionalOffers) ModifyPromotionalOffer() {}

// DeletePromotionalOfferFromSubscription Delete a specific promotional offer from an auto-renewable subscription.
func (cms *SubscriptionPromotionalOffers) DeletePromotionalOfferFromSubscription() {}
