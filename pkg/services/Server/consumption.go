package Server

import (
	"encoding/json"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/models"
	"net/http"
)

// ConsumptionService Consumption Information Service
type ConsumptionService struct {
	client *client.ServerClient
}

func NewConsumptionService(client *client.ServerClient) *ConsumptionService {
	return &ConsumptionService{
		client: client,
	}
}

// SendConsumptionInformation Send consumption information about a consumable in-app purchase or auto-renewable subscription to the App Store after your server receives a consumption request notification.
func (sis *ConsumptionService) SendConsumptionInformation(transactionId string, request *models.ConsumptionRequest) error {
	endpoint := fmt.Sprintf("inApps/v1/transactions/consumption/%s", transactionId)
	headers := map[string]string{
		"Accept": "application/json",
	}
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	_, code, err := sis.client.PUT(endpoint, headers, jsonBody)
	if err != nil {
		return err
	}
	switch code {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent,
		http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		return nil
	default:
		return fmt.Errorf("status code %d", code)
	}
}
