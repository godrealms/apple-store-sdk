package AppStore

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
)

type SubscriptionSubscriptionGroupSubmissions struct {
	client *client.AppStoreConnectAPIClient
}

func NewSubscriptionSubscriptionGroupSubmissions(client *client.AppStoreConnectAPIClient) *SubscriptionGroups {
	client.Config.BaseURL = "https://api.appstoreconnect.apple.com"
	return &SubscriptionGroups{client: client}
}

// CreateReviewSubmissionSubscriptionGroup Create a subscription group submission for review.
func (ssg *SubscriptionSubscriptionGroupSubmissions) CreateReviewSubmissionSubscriptionGroup() {}

// CreateReviewSubmissionSubscription Create a review submission for an auto-renewable subscription.
func (ssg *SubscriptionSubscriptionGroupSubmissions) CreateReviewSubmissionSubscription() {}
