package AppStore

import "github.com/godrealms/apple-store-sdk/pkg/client"

type AppBuildAndPrereleaseVersionInformation struct {
	client *client.AppStoreConnectAPIClient
}

func NewAppBuildAndPrereleaseVersionInformation(client *client.AppStoreConnectAPIClient) *AppBuildAndPrereleaseVersionInformation {
	return &AppBuildAndPrereleaseVersionInformation{client: client}
}
