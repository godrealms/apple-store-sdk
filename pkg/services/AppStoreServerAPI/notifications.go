package AppStoreServerAPI

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
	"net/http"
)

type NotificationService struct {
	client *client.ServerClient
}

func NewNotificationService(client *client.ServerClient) *NotificationService {
	return &NotificationService{
		client: client,
	}
}

// GetNotificationHistory Get a list of notifications that the App Store server attempted to send to your server.
// paginationToken: A pagination token that you return to the endpoint on a subsequent call to receive the next set of results.
func (ns *NotificationService) GetNotificationHistory(paginationToken string, request *models.NotificationHistoryRequest) (*models.NotificationHistoryResponse, error) {
	endpoint := "inApps/v1/notifications/history"
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body, code, err := ns.client.Post(endpoint, jsonBody, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.NotificationHistoryResponse
		if err = json.Unmarshal(body, &models.HistoryResponse{}); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// RequestTestNotification Ask App Store AppStoreServerAPI Notifications to send a test notification to your server.
func (ns *NotificationService) RequestTestNotification() (*models.SendTestNotificationResponse, error) {
	endpoint := "inApps/v1/notifications/test"
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ns.client.Post(endpoint, nil, headers)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.SendTestNotificationResponse
		if err = json.Unmarshal(body, &models.HistoryResponse{}); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}

// GetTestNotificationStatus Check the status of the test App Store server notification sent to your server.
func (ns *NotificationService) GetTestNotificationStatus(testNotificationToken string) (*models.CheckTestNotificationResponse, error) {
	endpoint := fmt.Sprintf("inApps/v1/notifications/test/%s", testNotificationToken)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ns.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		var response models.CheckTestNotificationResponse
		if err = json.Unmarshal(body, &response); err != nil {
			return nil, err
		}
		return &response, nil
	default:
		return nil, fmt.Errorf("status code %d", code)
	}
}
