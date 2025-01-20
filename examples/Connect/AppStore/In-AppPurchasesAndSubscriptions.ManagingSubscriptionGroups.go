package main

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/Connect/AppStore"
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

	appClient := client.NewConnectClient(config)
	service := AppStore.NewSubscriptionGroups(appClient)
	response, err := service.ListAllSubscriptionGroupsForAnApp(appid, nil)
	if err != nil {
		log.Fatalln("service.ListAllSubscriptionGroupsForAnApp failed: ", err)
		return
	}
	log.Printf("ListAllSubscriptionGroupsForAnApp response: %+v\n", response)
}
