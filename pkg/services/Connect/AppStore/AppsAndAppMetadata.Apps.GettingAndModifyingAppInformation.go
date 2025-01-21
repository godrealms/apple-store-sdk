package AppStore

import "github.com/godrealms/apple-store-sdk/pkg/client"

type GettingAndModifyingAppInformation struct {
	client *client.ConnectClient
}

func NewAppInformation(client *client.ConnectClient) *GettingAndModifyingAppInformation {
	return &GettingAndModifyingAppInformation{client: client}
}

// ListApps Find and list apps in App Store Connect.
func (g *GettingAndModifyingAppInformation) ListApps() {}

// ReadAppInformation Get information about a specific app.
func (g *GettingAndModifyingAppInformation) ReadAppInformation() {}

// ModifyAnApp Update app information, including bundle ID, primary locale, price schedule, and global availability.
func (g *GettingAndModifyingAppInformation) ModifyAnApp() {}

// ReadAnAppsEncryptionDeclarations Find and list all available app encryption declarations.
func (g *GettingAndModifyingAppInformation) ReadAnAppsEncryptionDeclarations() {}
