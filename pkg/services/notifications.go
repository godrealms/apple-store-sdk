package services

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
)

type NotificationService struct {
	client *client.Client
}

func NewNotificationService(client *client.Client) *NotificationService {
	return &NotificationService{
		client: client,
	}
}

// GetNotificationHistory Get a list of notifications that the App Store server attempted to send to your server.
// paginationToken: A pagination token that you return to the endpoint on a subsequent call to receive the next set of results.
func (ns *NotificationService) GetNotificationHistory(paginationToken string, request *models.NotificationHistoryRequest) (*models.NotificationHistoryResponse, error) {
	endpoint := "/inApps/v1/notifications/history"
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
	_ = body
	_ = code
	return nil, nil
}

// RequestTestNotification Ask App Store Server Notifications to send a test notification to your server.
func (ns *NotificationService) RequestTestNotification() (*models.SendTestNotificationResponse, error) {
	endpoint := "/inApps/v1/notifications/test"
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ns.client.Post(endpoint, nil, headers)
	if err != nil {
		return nil, err
	}
	_ = body
	_ = code
	return nil, nil
}

// GetTestNotificationStatus Check the status of the test App Store server notification sent to your server.
func (ns *NotificationService) GetTestNotificationStatus(testNotificationToken string) (*models.CheckTestNotificationResponse, error) {
	endpoint := fmt.Sprintf("/inApps/v1/notifications/test/%s", testNotificationToken)
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, code, err := ns.client.Get(endpoint, headers, nil)
	if err != nil {
		return nil, err
	}
	_ = body
	_ = code
	return nil, nil
}
