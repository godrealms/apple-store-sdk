package AppStore

import "github.com/godrealms/apple-store-sdk/pkg/client"

type AppBuildAndPrereleaseVersionInformation struct {
	client *client.ConnectClient
}

func NewAppBuildAndPrereleaseVersionInformation(client *client.ConnectClient) *AppBuildAndPrereleaseVersionInformation {
	return &AppBuildAndPrereleaseVersionInformation{client: client}
}
