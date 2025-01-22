package main

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/AppStoreConnectAPI/AppStore"
	"log"
)

func main() {
	Iss := ""
	Bid := ""
	Kid := ""
	PrivateKey := ""
	appid := ""
	//groupId := ""
	config, err := client.NewConfig(true, Iss, Bid, Kid, PrivateKey)
	if err != nil {
		log.Fatalln(err)
		return
	}

	appClient := client.NewAppStoreConnectAPIClient(config)
	service := AppStore.NewSubscriptionGroups(appClient)
	response, err := service.ListAllSubscriptionGroupLocalizations(appid, nil)
	if err != nil {
		log.Fatalln("service.ListAllSubscriptionGroupLocalizations failed: ", err)
		return
	}
	log.Printf("ListAllSubscriptionGroupLocalizations response: %+v\n", response)
	//response, err := service.ListAllSubscriptionGroupsForAnApp(appid, nil)
	//if err != nil {
	//	log.Fatalln("service.ListAllSubscriptionGroupsForAnApp failed: ", err)
	//	return
	//}
	//log.Printf("ListAllSubscriptionGroupsForAnApp response: %+v\n", response)
}
