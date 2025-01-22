package main

import (
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreServerNotifications/ServerNotificationsVersion2"
	"log"
)

func main() {
	signedPayload := ""
	notifications, err := serverNotificationsVersion2.Notifications(signedPayload)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%+v\n", notifications)
}
