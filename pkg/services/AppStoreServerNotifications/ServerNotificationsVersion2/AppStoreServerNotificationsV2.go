package serverNotificationsVersion2

import "github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerNotifications/ResponseTypes"

type AppStoreServerNotificationsV2 struct{}

func NewAppStoreServerNotificationsV2() *AppStoreServerNotificationsV2 {
	return &AppStoreServerNotificationsV2{}
}

func Notifications(signedPayload string) (*responseTypes.ResponseBodyV2DecodedPayload, error) {
	notificationsResponseBodyV2 := responseTypes.NotificationsResponseBodyV2{
		SignedPayload: responseTypes.SignedPayload(signedPayload),
	}
	return notificationsResponseBodyV2.SignedPayload.DecodedPayload()
}
