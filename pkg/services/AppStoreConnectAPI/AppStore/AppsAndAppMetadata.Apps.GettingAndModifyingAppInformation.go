package AppStore

import "github.com/godrealms/apple-store-sdk/pkg/client"

type GettingAndModifyingAppInformation struct {
	client *client.AppStoreConnectAPIClient
}

func NewAppInformation(client *client.AppStoreConnectAPIClient) *GettingAndModifyingAppInformation {
	return &GettingAndModifyingAppInformation{client: client}
}

// ListApps Find and list apps in App Store AppStoreConnectAPI.
func (g *GettingAndModifyingAppInformation) ListApps() {}

// ReadAppInformation Get information about a specific app.
func (g *GettingAndModifyingAppInformation) ReadAppInformation() {}

// ModifyAnApp Update app information, including bundle ID, primary locale, price schedule, and global availability.
func (g *GettingAndModifyingAppInformation) ModifyAnApp() {}

// ReadAnAppsEncryptionDeclarations Find and list all available app encryption declarations.
func (g *GettingAndModifyingAppInformation) ReadAnAppsEncryptionDeclarations() {}
