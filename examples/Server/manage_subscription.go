package main

import (
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"github.com/godrealms/apple-store-sdk/pkg/services/Server"
	"log"
)

func main() {
	Iss := ""           // Your issuer ID from the Keys page in App Store Connect (Ex: “57246542-96fe-1a63-e053-0824d011072a")
	Bid := ""           // Your app’s bundle ID (Ex: “com.example.testbundleid”)
	Kid := ""           // Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	PrivateKey := ""    // Private key string corresponding to the private key ID from App Store Connect
	transactionId := "" // (Required) The identifier of a transaction that belongs to the customer, and which may be an original transaction identifier (originalTransactionId).                                                                                                                                                                                 //
	config, err := client.NewConfig(true, Iss, Bid, Kid, PrivateKey)
	if err != nil {
		log.Fatalln(err)
		return
	}

	appClient := client.NewServerClient(config)
	service := Server.NewSubscriptionsService(appClient)

	SubscriptionStatuses, err := service.GetAllSubscriptionStatuses(transactionId)
	if err != nil {
		log.Fatalf("GetAllSubscriptionStatuses err: %+v\n", err)
		return
	}
	//log.Printf("SubscriptionStatuses: %+v\n", SubscriptionStatuses)

	for _, item := range SubscriptionStatuses.Data {
		log.Printf("SubscriptionGroupIdentifier: %+v\n", item.SubscriptionGroupIdentifier)
		for _, transaction := range item.LastTransactions {
			log.Printf("OriginalTransactionId: %+v\n", transaction.OriginalTransactionId)
			log.Printf("Status: %+v\n", transaction.Status)
			decrypt, err := transaction.SignedRenewalInfo.Decrypt()
			if err != nil {
				log.Printf("transaction.SignedRenewalInfo.Decrypt: %+v\n", err)
			}
			log.Printf("SignedRenewalInfo: %+v\n", decrypt)
			payload, err := transaction.SignedTransactionInfo.Decrypt()
			if err != nil {
				log.Printf("transaction.SignedTransactionInfo.Decrypt: %+v\n", err)
			}
			log.Printf("SignedTransactionInfo: %+v\n", payload)
		}
	}
}
