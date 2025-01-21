package TypeAliases

// BuildAudienceType A string that represents the App Store AppStoreConnectAPI audience for a build.
type BuildAudienceType string

const (
	BuildAudienceTypeInternalOnly     BuildAudienceType = "INTERNAL_ONLY"      // The build of your app is only available to members of your development team.
	BuildAudienceTypeAppStoreEligible BuildAudienceType = "APP_STORE_ELIGIBLE" // The build of your app is eligible for submission and release on the App Store.
)

func (t BuildAudienceType) String() string {
	return string(t)
}
