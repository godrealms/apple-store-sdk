package AppStore

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreConnectAPI/TypeAliases"
	"net/http"
)

type SubscriptionGroups struct {
	client *client.ConnectClient
}

func NewSubscriptionGroups(client *client.ConnectClient) *SubscriptionGroups {
	client.Config.BaseURL = "https://api.appstoreconnect.apple.com"
	return &SubscriptionGroups{client: client}
}

// CreateSubscriptionGroup Create a subscription group for an app.
func (sg *SubscriptionGroups) CreateSubscriptionGroup() {}

// ListAllSubscriptionGroupsForAnApp Get a list of subscription groups for a specific app.
func (sg *SubscriptionGroups) ListAllSubscriptionGroupsForAnApp(appid string, parameters *TypeAliases.SubscriptionGroupsQueryParameters) (*TypeAliases.SubscriptionGroupsResponse, error) {
	endpoint := fmt.Sprintf("/v1/apps/%s/subscriptionGroups", appid)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := sg.client.Get(endpoint, headers, parameters)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		response := TypeAliases.SubscriptionGroupsResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// ReadSubscriptionGroupInformation Get the details of a specific subscription group.
func (sg *SubscriptionGroups) ReadSubscriptionGroupInformation() {}

// ModifySubscriptionGroup Update the reference name for a specific subscription group.
func (sg *SubscriptionGroups) ModifySubscriptionGroup() {}

// DeleteSubscriptionGroup Delete a specific empty subscription group.
func (sg *SubscriptionGroups) DeleteSubscriptionGroup() {}

// ListAllSubscriptionGroupLocalizations Get a list of all localized metadata for a specific subscription group.
func (sg *SubscriptionGroups) ListAllSubscriptionGroupLocalizations(appid string, parameters *TypeAliases.SubscriptionGroupsQueryParameters) (*TypeAliases.SubscriptionGroupLocalizationsResponse, error) {
	endpoint := fmt.Sprintf("v1/subscriptionGroups/%s/subscriptionGroupLocalizations", appid)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := sg.client.Get(endpoint, headers, parameters)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		response := TypeAliases.SubscriptionGroupLocalizationsResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// ListAllSubscriptionsForSubscriptionGroup Get a list of all auto-renewable subscriptions in a subscription group.
func (sg *SubscriptionGroups) ListAllSubscriptionsForSubscriptionGroup() {}
