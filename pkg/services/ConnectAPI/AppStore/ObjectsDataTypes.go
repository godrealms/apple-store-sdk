package AppStore

type uri string
type uriReference string

// SubscriptionStatusUrlVersion Strings that represent versions of App Store Server Notifications.
// V1: Represents version 1 of App Store Server Notifications,
// and indicates you receive notifications on your server at the App Store Server Notifications V1 endpoint.
// Note that App Store Server Notifications V1 are deprecated.
//
// V2: Represents version 2 of App Store Server Notifications,
// and indicates you receive notifications on your server at the App Store Server Notifications V2 endpoint.
type SubscriptionStatusUrlVersion string

// Attributes that describe an Apps resource.
type Attributes struct {
	// The bundle ID for your app. This ID must match the one you use in Xcode.
	// The bundle ID cannot be changed after you upload your first build.
	BundleId string `json:"bundleId"`

	// The name of your app as it will appear in the App Store. The maximum length is 30 characters.
	Name string `json:"name"`

	// The primary locale for your app. If localized app information isn’t available in an App Store territory,
	// the information from your primary language is used instead.
	PrimaryLocale string `json:"primaryLocale"`

	// A unique ID for your app that is not visible on the App Store.
	Sku string `json:"sku"`

	// Possible Values: DOES_NOT_USE_THIRD_PARTY_CONTENT, USES_THIRD_PARTY_CONTENT
	ContentRightsDeclaration string `json:"contentRightsDeclaration"`

	IsOrEverWasMadeForKids bool `json:"isOrEverWasMadeForKids"`

	SubscriptionStatusUrl uri `json:"subscriptionStatusUrl"`

	SubscriptionStatusUrlForSandbox uri `json:"subscriptionStatusUrlForSandbox"`

	SubscriptionStatusUrlVersion SubscriptionStatusUrlVersion `json:"subscriptionStatusUrlVersion"`

	SubscriptionStatusUrlVersionForSandbox SubscriptionStatusUrlVersion `json:"subscriptionStatusUrlVersionForSandbox"`
}

// Data The type and ID of a related resource.
type Data struct {
	// (Required) The opaque resource ID that uniquely identifies the resource.
	ID string `json:"id"`

	// (Required) The resource type.
	// Value: betaLicenseAgreements
	// Value: betaAppReviewDetails
	Type string `json:"type"`
}

// RelationshipLinks Links related to the response document, including self links.
type RelationshipLinks struct {
	// The link to the related documents.
	Related uriReference `json:"related"`
	// The link that produced the current document.
	Self uriReference `json:"self"`
}

type DataLinks struct {
	Data  Data              `json:"data"`
	Links RelationshipLinks `json:"links"`
}

type DataLinksMeta struct {
	Data  []Data            `json:"data"`
	Links RelationshipLinks `json:"links"`
	Meta  PagingInformation `json:"meta"`
}

// BetaLicenseAgreement The data and links that describe the relationship between the resources.
type BetaLicenseAgreement DataLinks

// Paging details such as the total number of resources and the per-page limit.
type Paging struct {
	//The total number of resources matching your request.
	Total int32 `json:"total"`
	// (Required) The maximum number of resources to return per page, from 0 to 200.
	Limit int32 `json:"limit"`
}

// PagingInformation Paging information for data responses.
type PagingInformation struct {
	// (Required) The paging information details.
	Paging Paging `json:"paging"`
}

// PreReleaseVersions The data and links that describe the relationship between the resources.
type PreReleaseVersions DataLinksMeta

// BetaAppLocalizations The data and links that describe the relationship between the resources.
type BetaAppLocalizations DataLinksMeta

// BetaGroups The data and links that describe the relationship between the resources.
type BetaGroups DataLinksMeta

// Builds The data and links that describe the relationship between the resources.
type Builds DataLinksMeta

// BetaAppReviewDetail The data and links that describe the relationship between the Apps and the Beta App Review Details resources.
type BetaAppReviewDetail DataLinks

// AppInfos The data and links that describe the relationship between the resources.
type AppInfos DataLinksMeta

// Versions The data and links that describe the relationship between the resources.
type Versions DataLinksMeta

// EndUserLicenseAgreement The data and links that describe the relationship between the resources.
type EndUserLicenseAgreement DataLinks

// GameCenterEnabledVersions The data and links that describe the relationship between the resources.
type GameCenterEnabledVersions DataLinksMeta

// InAppPurchases The data and links that describe the relationship between the resources.
type InAppPurchases DataLinksMeta

// CiProduct The data and links that describe the relationship between the Apps and Products resources.
type CiProduct DataLinks

// AppClips The data and links that describe the relationship between the Apps and the App Clips resources.
type AppClips DataLinksMeta

type AppCustomProductPages DataLinksMeta
type AppEvents DataLinksMeta
type ReviewSubmissions DataLinksMeta
type SubscriptionGracePeriod DataLinks
type InAppPurchasesV2 DataLinksMeta
type PromotedPurchases DataLinksMeta
type SubscriptionGroups DataLinksMeta
type VersionExperimentsV2 DataLinksMeta
type AppEncryptionDeclarations DataLinksMeta
type GameCenterDetail DataLinks
type AlternativeDistributionKey struct {
	Links RelationshipLinks `json:"links"`
}
type AnalyticsReportRequests struct {
	Links RelationshipLinks `json:"links"`
}
type AppAvailabilityV2 struct {
	Links RelationshipLinks `json:"links"`
}
type AppPricePoints struct {
	Links RelationshipLinks `json:"links"`
}
type AppPriceSchedule struct {
	Links RelationshipLinks `json:"links"`
}
type BetaTesters struct {
	Links RelationshipLinks `json:"links"`
}
type CustomerReviews struct {
	Links RelationshipLinks `json:"links"`
}
type MarketplaceSearchDetail struct {
	Links RelationshipLinks `json:"links"`
}
type PerfPowerMetrics struct {
	Links RelationshipLinks `json:"links"`
}

// Relationships The relationships you included in the request and those on which you can operate.
type Relationships struct {
	// The data and links that describe the relationship between the Apps and the Beta License Agreements resources.
	BetaLicenseAgreement BetaLicenseAgreement `json:"betaLicenseAgreement"`

	// The data and links that describe the relationship between the Apps and the Pre-Release Versions resources.
	PreReleaseVersions PreReleaseVersions `json:"preReleaseVersions"`

	// The data and links that describe the relationship between the Apps and the Beta App Localizations resources.
	BetaAppLocalizations BetaAppLocalizations `json:"betaAppLocalizations"`

	// The data and links that describe the relationship between the Apps and the Beta Groups resources.
	BetaGroups BetaGroups `json:"betaGroups"`

	// The data and links that describe the relationship between the Apps and the Builds resources.
	Builds Builds `json:"builds"`

	// The data and links that describe the relationship between the Apps and the Beta App Review Details resources.
	BetaAppReviewDetail BetaAppReviewDetail `json:"betaAppReviewDetail"`

	// The data and links that describe the relationship between the Apps and the App Infos resources.
	AppInfos AppInfos `json:"appInfos"`

	// The data and links that describe the relationship between the Apps and the App Store Versions resources.
	AppStoreVersions Versions `json:"appStoreVersions"`

	// The data and links that describe the relationship between the Apps and the End User License Agreements (EULA) resources.
	EndUserLicenseAgreement EndUserLicenseAgreement `json:"endUserLicenseAgreement"`

	// The data and links that describe the relationship between the Apps and the Game Center Enabled Versions resources.
	GameCenterEnabledVersions GameCenterEnabledVersions `json:"gameCenterEnabledVersions"`

	// Deprecated  The data and links that describe the relationship between the Apps and the In App Purchases resources.
	InAppPurchases InAppPurchases `json:"inAppPurchases"`

	// The data and links that describe the relationship between the Apps and the Products resources.
	CiProduct CiProduct `json:"ciProduct"`

	// The data and links that describe the relationship between the Apps and the App Clips resources.
	AppClips AppClips `json:"appClips"`

	AppCustomProductPages        AppCustomProductPages      `json:"appCustomProductPages"`
	AppEvents                    AppEvents                  `json:"appEvents"`
	ReviewSubmissions            ReviewSubmissions          `json:"reviewSubmissions"`
	SubscriptionGracePeriod      SubscriptionGracePeriod    `json:"subscriptionGracePeriod"`
	InAppPurchasesV2             InAppPurchasesV2           `json:"inAppPurchasesV2"`
	PromotedPurchases            PromotedPurchases          `json:"promotedPurchases"`
	SubscriptionGroups           SubscriptionGroups         `json:"subscriptionGroups"`
	AppStoreVersionExperimentsV2 VersionExperimentsV2       `json:"appStoreVersionExperimentsV2"`
	AppEncryptionDeclarations    AppEncryptionDeclarations  `json:"appEncryptionDeclarations"`
	GameCenterDetail             GameCenterDetail           `json:"gameCenterDetail"`
	AlternativeDistributionKey   AlternativeDistributionKey `json:"alternativeDistributionKey"`
	AnalyticsReportRequests      AnalyticsReportRequests    `json:"analyticsReportRequests"`
	AppAvailabilityV2            AppAvailabilityV2          `json:"appAvailabilityV2"`
	AppPricePoints               AppPricePoints             `json:"appPricePoints"`
	AppPriceSchedule             AppPriceSchedule           `json:"appPriceSchedule"`
	BetaTesters                  BetaTesters                `json:"betaTesters"`
	CustomerReviews              CustomerReviews            `json:"customerReviews"`
	MarketplaceSearchDetail      MarketplaceSearchDetail    `json:"marketplaceSearchDetail"`
	PerfPowerMetrics             PerfPowerMetrics           `json:"perfPowerMetrics"`
}

// ResourceLinks Self-links to requested resources.
type ResourceLinks struct {
	// The link to the resource.
	Self uriReference `json:"self"`
}

// DocumentLinks (Required) The link that produced the current document.
type DocumentLinks struct {
	// The link to the resource.
	Self uriReference `json:"self"`
}

// App The data structure that represents an Apps resource.
type App struct {
	// (Required) The opaque resource ID that uniquely identifies the resource.
	ID string `json:"id"`

	// The resource’s attributes.
	Attributes Attributes `json:"attributes"`

	// Navigational links to related data and included resource types and IDs.
	Relationships Relationships `json:"relationships"`

	// (Required) The resource type.
	Type string `json:"type"`

	// Navigational links that include the self-link.
	Links ResourceLinks `json:"links"`
}

type QueryParameters struct {
	// Fields to return for included related types.
	// fields[apps][string]
	// Possible Values: name, bundleId, sku, primaryLocale, isOrEverWasMadeForKids, subscriptionStatusUrl,
	// subscriptionStatusUrlVersion, subscriptionStatusUrlForSandbox, subscriptionStatusUrlVersionForSandbox,
	// contentRightsDeclaration, streamlinedPurchasingEnabled, appEncryptionDeclarations, ciProduct, betaTesters,
	// betaGroups, appStoreVersions, preReleaseVersions, betaAppLocalizations, builds, betaLicenseAgreement,
	// betaAppReviewDetail, appInfos, appClips, appPricePoints, endUserLicenseAgreement, appPriceSchedule,
	// appAvailabilityV2, inAppPurchases, subscriptionGroups, gameCenterEnabledVersions, perfPowerMetrics,
	// appCustomProductPages, inAppPurchasesV2, promotedPurchases, appEvents, reviewSubmissions, subscriptionGracePeriod,
	// customerReviews, gameCenterDetail, appStoreVersionExperimentsV2, alternativeDistributionKey,
	// analyticsReportRequests, marketplaceSearchDetail
	//
	// fields[betaLicenseAgreements][string]
	// Possible Values: agreementText, app
	//
	// fields[preReleaseVersions][string]
	// Possible Values: version, platform, builds, app
	//
	// fields[betaAppLocalizations][string]
	// Possible Values: feedbackEmail, marketingUrl, privacyPolicyUrl, tvOsPrivacyPolicy, description, locale, app
	//
	// fields[builds][string]
	// Possible Values: version, uploadedDate, expirationDate, expired, minOsVersion, lsMinimumSystemVersion,
	// computedMinMacOsVersion, iconAssetToken, processingState, buildAudienceType, usesNonExemptEncryption,
	// preReleaseVersion, individualTesters, betaGroups, betaBuildLocalizations, appEncryptionDeclaration,
	// betaAppReviewSubmission, app, buildBetaDetail, appStoreVersion, icons, buildBundles, perfPowerMetrics, diagnosticSignatures
	//
	// fields[betaGroups][string]
	// Possible Values: name, createdDate, isInternalGroup, hasAccessToAllBuilds, publicLinkEnabled, publicLinkId,
	// publicLinkLimitEnabled, publicLinkLimit, publicLink, feedbackEnabled, iosBuildsAvailableForAppleSiliconMac, app,
	// builds, betaTesters
	//
	// fields[endUserLicenseAgreements][string]
	// Possible Values: agreementText, app, territories
	//
	// Attributes, relationships, and IDs by which to filter.
	// filter[bundleId][string]
	// filter[id][string]
	// filter[name][string]
	// filter[sku][string]
	//
	// Fields to return for included related types. Note: appStoreState is deprecated, use appVersionState instead.
	// fields[appStoreVersions][string]
	// Possible Values: platform, versionString, appStoreState, appVersionState, copyright, reviewType, releaseType,
	// earliestReleaseDate, downloadable, createdDate, app, ageRatingDeclaration, appStoreVersionLocalizations, build,
	// appStoreVersionPhasedRelease, gameCenterAppVersion, routingAppCoverage, appStoreReviewDetail, appStoreVersionSubmission,
	// appClipDefaultExperience, appStoreVersionExperiments, appStoreVersionExperimentsV2, customerReviews, alternativeDistributionPackage
	//
	// fields[appInfos][string]
	// Possible Values: appStoreState, state, appStoreAgeRating, australiaAgeRating, brazilAgeRating, brazilAgeRatingV2,
	// franceAgeRating, koreaAgeRating, kidsAgeBand, app, ageRatingDeclaration, appInfoLocalizations, primaryCategory,
	// primarySubcategoryOne, primarySubcategoryTwo, secondaryCategory, secondarySubcategoryOne, secondarySubcategoryTwo
	//
	// filter[appStoreVersions][string]
	//
	// filter[appStoreVersions.platform][string]
	// Possible Values: IOS, MAC_OS, TV_OS, VISION_OS
	//
	// filter[appStoreVersions.appStoreState][string]
	// Possible Values: ACCEPTED, DEVELOPER_REMOVED_FROM_SALE, DEVELOPER_REJECTED, IN_REVIEW, INVALID_BINARY,
	// METADATA_REJECTED, PENDING_APPLE_RELEASE, PENDING_CONTRACT, PENDING_DEVELOPER_RELEASE, PREPARE_FOR_SUBMISSION,
	// PREORDER_READY_FOR_SALE, PROCESSING_FOR_APP_STORE, READY_FOR_REVIEW, READY_FOR_SALE, REJECTED, REMOVED_FROM_SALE,
	// WAITING_FOR_EXPORT_COMPLIANCE, WAITING_FOR_REVIEW, REPLACED_WITH_NEW_VERSION, NOT_APPLICABLE
	//
	// fields[inAppPurchases][string]
	// Possible Values: referenceName, productId, inAppPurchaseType, state, apps, name, reviewNote, familySharable,
	// contentHosting, inAppPurchaseLocalizations, pricePoints, content, appStoreReviewScreenshot, promotedPurchase,
	// iapPriceSchedule, inAppPurchaseAvailability, images
	//
	// fields[ciProducts][string]
	// Possible Values: name, createdDate, productType, app, bundleId, workflows, primaryRepositories, additionalRepositories, buildRuns
	// fields[appClips][string]
	// Possible Values: bundleId, app, appClipDefaultExperiences, appClipAdvancedExperiences
	//
	// fields[reviewSubmissions] [string]
	// Possible Values: platform, submittedDate, state, app, items, appStoreVersionForReview, submittedByActor, lastUpdatedByActor
	//
	// fields[appCustomProductPages][string]
	// Possible Values: name, url, visible, app, appCustomProductPageVersions
	//
	// fields[appEvents][string]
	// Possible Values: referenceName, badge, eventState, deepLink, purchaseRequirement, primaryLocale, priority, purpose,
	// territorySchedules, archivedTerritorySchedules, localizations
	//
	// fields[subscriptionGracePeriods][string]
	// Possible Values: optIn, sandboxOptIn, duration, renewalType
	//
	// fields[promotedPurchases][string]
	// Possible Values: visibleForAllUsers, enabled, state, inAppPurchaseV2, subscription, promotionImages
	//
	// fields[subscriptionGroups][string]
	// Possible Values: referenceName, subscriptions, subscriptionGroupLocalizations
	//
	// fields[appStoreVersionExperiments][string]
	// Possible Values: name, platform, trafficProportion, state, reviewRequired, startDate, endDate, app,
	// latestControlVersion, controlVersions, appStoreVersionExperimentTreatments
	//
	// fields[appEncryptionDeclarations][string]
	// Possible Values: appDescription, createdDate, usesEncryption, exempt, containsProprietaryCryptography,
	// containsThirdPartyCryptography, availableOnFrenchStore, platform, uploadedDate, documentUrl, documentName,
	// documentType, appEncryptionDeclarationState, codeValue, app, builds, appEncryptionDeclarationDocument
	//
	// fields[gameCenterDetails][string]
	// Possible Values: arcadeEnabled, challengeEnabled, app, gameCenterAppVersions, gameCenterGroup, gameCenterLeaderboards,
	// gameCenterLeaderboardSets, gameCenterAchievements, defaultLeaderboard, defaultGroupLeaderboard, achievementReleases,
	// leaderboardReleases, leaderboardSetReleases
	//
	// This filter is deprecated.
	// filter[appStoreVersions.appVersionState][string]
	// Possible Values: ACCEPTED, DEVELOPER_REJECTED, IN_REVIEW, INVALID_BINARY, METADATA_REJECTED, PENDING_APPLE_RELEASE,
	// PENDING_DEVELOPER_RELEASE, PREPARE_FOR_SUBMISSION, PROCESSING_FOR_DISTRIBUTION, READY_FOR_DISTRIBUTION,
	// READY_FOR_REVIEW, REJECTED, REPLACED_WITH_NEW_VERSION, WAITING_FOR_EXPORT_COMPLIANCE, WAITING_FOR_REVIEW
	//
	// filter[reviewSubmissions.platform][string]
	// Possible Values: IOS, MAC_OS, TV_OS, VISION_OS
	//
	// filter[reviewSubmissions.state][string]
	// Possible Values: READY_FOR_REVIEW, WAITING_FOR_REVIEW, IN_REVIEW, UNRESOLVED_ISSUES, CANCELING, COMPLETING, COMPLETE
	Fields []string `json:"fields"`

	// Relationship data to include in the response.
	// Possible Values: appEncryptionDeclarations, ciProduct, betaGroups, appStoreVersions, preReleaseVersions,
	// betaAppLocalizations, builds, betaLicenseAgreement, betaAppReviewDetail, appInfos, appClips, endUserLicenseAgreement,
	// inAppPurchases, subscriptionGroups, gameCenterEnabledVersions, appCustomProductPages, inAppPurchasesV2,
	// promotedPurchases, appEvents, reviewSubmissions, subscriptionGracePeriod, gameCenterDetail, appStoreVersionExperimentsV2
	Include []string `json:"include"`

	// Number of included related resources to return.
	// Number of resources to return.
	// Maximum: 200
	//
	// limit[builds]integer
	// limit[betaGroups]integer
	// limit[betaAppLocalizations]integer
	// limit[appStoreVersions]integer
	// limit[appInfos]integer
	// limit[appClips]integer
	// limit[appCustomProductPages]integer
	// limit[appEvents]integer
	// limit[reviewSubmissions]integer
	// limit[inAppPurchasesV2]integer
	// limit[promotedPurchases]integer
	// limit[subscriptionGroups]integer
	// limit[appStoreVersionExperimentsV2]integer
	// limit[appEncryptionDeclarations]integer
	// Maximum: 50
	Limit int32 `json:"limit"`

	// Attributes by which to sort.
	// Possible Values: name, -name, bundleId, -bundleId, sku, -sku
	Sort []string `json:"sort"`
}

// PagedDocumentLinks Links related to the response document, including paging links.
type PagedDocumentLinks struct {
	// The link to the first page of documents.
	First uriReference `json:"first"`
	// The link to the next page of documents.
	Next uriReference `json:"next"`
	// (Required) The link that produced the current document.
	Self uriReference `json:"self"`
}

// AppsResponse A response that contains a list of Apps resources.
type AppsResponse struct {
	// (Required) The resource data.
	Data []App `json:"data"`

	// (Required) Navigational links that include the self-link.
	Links PagedDocumentLinks `json:"links"`

	// Paging information.
	Meta PagingInformation `json:"meta"`

	// Possible types: AppEncryptionDeclaration, CiProduct, BetaGroup, AppStoreVersion, PrereleaseVersion,
	// BetaAppLocalization, Build, BetaLicenseAgreement, BetaAppReviewDetail, AppInfo, AppClip, EndUserLicenseAgreement,
	// InAppPurchase, SubscriptionGroup, GameCenterEnabledVersion, AppCustomProductPage, InAppPurchaseV2, PromotedPurchase,
	// AppEvent, ReviewSubmission, SubscriptionGracePeriod, GameCenterDetail, AppStoreVersionExperimentV2
	Included []any `json:"included"`
}

// AppResponse A response that contains a single Apps resource.
type AppResponse struct {
	// (Required) The resource data.
	Data App `json:"data"`

	// (Required) Navigational links that include the self-link.
	Links DocumentLinks `json:"links"`

	// Possible types: AppEncryptionDeclaration, CiProduct, BetaGroup, AppStoreVersion, PrereleaseVersion,
	// BetaAppLocalization, Build, BetaLicenseAgreement, BetaAppReviewDetail, AppInfo, AppClip, EndUserLicenseAgreement,
	// InAppPurchase, SubscriptionGroup, GameCenterEnabledVersion, AppCustomProductPage, InAppPurchaseV2, PromotedPurchase,
	// AppEvent, ReviewSubmission, SubscriptionGracePeriod, GameCenterDetail, AppStoreVersionExperimentV2
	Included []any `json:"included"`
}

type AppUpdateRequest struct{}
type AppEncryptionDeclarationsResponse string
type BuildsWithoutIncludesResponse string
type PreReleaseVersionsWithoutIncludesResponse string
type AppClipsResponse string
type BetaGroupsWithoutIncludesResponse string
type BetaAppReviewDetailWithoutIncludesResponse string
type BetaLicenseAgreementWithoutIncludesResponse string
type CiProductResponse string
type AppPricePointsV3Response string
type AppInfosResponse string
type VersionsResponse string
type EndUserLicenseAgreementWithoutIncludesResponse string
type AppCustomProductPagesResponse string
type VersionExperimentsV2Response string
type PromotedPurchasesResponse string
type ReviewSubmissionsResponse string
type XcodeMetrics string
type CustomerReviewsResponse string
type TerritoryResponse string
type AppPricesV2Response string
type AppPriceScheduleResponse string
type AppAvailabilityV2Response string
type AlternativeDistributionKeyResponse string
type MarketplaceSearchDetailResponse string
type AppPriceScheduleCreateRequest string
