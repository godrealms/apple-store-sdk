package server

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
	"io/ioutil"
	"net/http"
)

type NotificationService struct {
	client *client.Client
}

func NewNotificationService(client *client.Client) *NotificationService {
	return &NotificationService{
		client: client,
	}
}

// ReceiveNotifications Receiving App Store Server Notifications
func (ns *NotificationService) ReceiveNotifications(request *http.Request) (*models.NotificationsResponseBodyV2DecodedPayload, error) {
	// 读取请求体
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %w", err)
	}
	defer request.Body.Close() // 确保请求体关闭

	// 解析 JSON 数据
	var requestData models.NotificationsResponseBodyV2
	if err = json.Unmarshal(body, &requestData); err != nil {
		return nil, fmt.Errorf("error unmarshalling request body: %w", err)
	}

	return requestData.SignedPayload.DecodedPayload()
}
