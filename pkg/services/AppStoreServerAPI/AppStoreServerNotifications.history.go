package appStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/RequestTypes"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerAPI/ResponseTypes"
	"github.com/godrealms/apple-store-sdk/pkg/types"
	"net/http"
)

type AppStoreServerNotificationsHistory struct {
	client *client.AppStoreServerAPIClient
}

// NewAppStoreServerNotificationsHistory
// Get a list of notifications that the App Store server attempted to send to your server.
func NewAppStoreServerNotificationsHistory(client *client.AppStoreServerAPIClient) *AppStoreServerNotificationsHistory {
	return &AppStoreServerNotificationsHistory{client: client}
}

// GetNotificationHistory Get a list of notifications that the App Store server attempted to send to your server.
// paginationToken: A pagination token that you return to the endpoint on a subsequent call to receive the next set of results.
func (snh *AppStoreServerNotificationsHistory) GetNotificationHistory(paginationToken types.PaginationToken, request *RequestTypes.NotificationHistoryRequest) (*ResponseTypes.NotificationHistoryResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/notifications/history%s", paginationToken.GetParameter())
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := snh.client.Post(endpoint, jsonBody, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.NotificationHistoryResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// RequestTestNotification Ask App Store AppStoreServerAPI Notifications to send a test notification to your server.
func (snh *AppStoreServerNotificationsHistory) RequestTestNotification() (*ResponseTypes.SendTestNotificationResponse, error) {
	endpoint := "/inApps/v1/notifications/test"
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := snh.client.Post(endpoint, nil, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.SendTestNotificationResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// GetTestNotificationStatus Check the status of the test App Store server notification sent to your server.
func (snh *AppStoreServerNotificationsHistory) GetTestNotificationStatus(testNotificationToken string) (*ResponseTypes.CheckTestNotificationResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/notifications/test/%s", testNotificationToken)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := snh.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response ResponseTypes.CheckTestNotificationResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
