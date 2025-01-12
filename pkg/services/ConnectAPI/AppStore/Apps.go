package AppStore

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/godrealms/apple-store-sdk/pkg/client"
	"net/http"
	"strconv"
)

type Apps struct {
	client *client.Client
}

// NewApps
// An apps resource represents your app that’s currently in development,
// or available on the App Store through the App Store Connect website.
// Use the apps resource to manage and maintain your existing apps.
//
// Don’t use this API to create new apps; instead, create new apps on the App Store Connect website.
// To upload builds to App Store Connect, you must use Xcode, Transporter, or the Transporter Mac app.
// This API doesn’t permit you to directly upload your builds,
// but you may use App Store Connect API Keys in conjunction with Transporter to upload.
// To download the Transporter app, see the Mac App Store.
//
// To learn more about managing your apps, see Add a new app.
func NewApps(client *client.Client) *Apps {
	client.Config.SetBaseUrl("https://api.appstoreconnect.apple.com")
	return &Apps{client: client}
}

// ListApps Find and list apps in App Store Connect.
func (a *Apps) ListApps(parameters *QueryParameters) (*AppsResponse, error) {
	url := fmt.Sprintf("/v1/apps")
	headers := map[string]string{}
	body, status, err := a.client.Get(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var apps AppsResponse
	if err = json.Unmarshal(body, &apps); err != nil {
		return nil, err
	}
	return &apps, nil
}

// ReadAppInformation Get information about a specific app.
func (a *Apps) ReadAppInformation(id string, parameters *QueryParameters) (*AppResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s", id)
	headers := map[string]string{}
	body, status, err := a.client.Get(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var apps AppResponse
	if err = json.Unmarshal(body, &apps); err != nil {
		return nil, err
	}
	return &apps, nil
}

// ModifyAnApp Update app information, including bundle ID, primary locale, price schedule, and global availability.
func (a *Apps) ModifyAnApp(id string, parameters *AppUpdateRequest) (*AppResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var apps AppResponse
	if err = json.Unmarshal(body, &apps); err != nil {
		return nil, err
	}
	return &apps, nil
}

// ReadAnAppsEncryptionDeclarations Find and list all available app encryption declarations.
func (a *Apps) ReadAnAppsEncryptionDeclarations(id string, parameters *QueryParameters) (*AppEncryptionDeclarationsResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appEncryptionDeclarations", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppEncryptionDeclarationsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllBuildsOfAnApp Get a list of builds associated with a specific app.
func (a *Apps) ListAllBuildsOfAnApp(id string, parameters *QueryParameters) (*BuildsWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/builds", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response BuildsWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllPrereleaseVersionsForAnApp Get a list of prerelease versions associated with a specific app.
func (a *Apps) ListAllPrereleaseVersionsForAnApp(id string, parameters *QueryParameters) (*PreReleaseVersionsWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/preReleaseVersions", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response PreReleaseVersionsWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllAppClipsForAnApp List your app’s associated App Clips.
func (a *Apps) ListAllAppClipsForAnApp(id string, parameters *QueryParameters) (*AppClipsResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appClips", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppClipsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllBetaGroupsForAnApp Get a list of beta groups associated with a specific app.
func (a *Apps) ListAllBetaGroupsForAnApp(id string, parameters *QueryParameters) (*BetaGroupsWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/betaGroups", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response BetaGroupsWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RemoveSpecifiedBetaTestersFromAllGroupsAndBuildsOfAnApp Remove one or more beta testers’ access to test any builds of a specific app.
func (a *Apps) RemoveSpecifiedBetaTestersFromAllGroupsAndBuildsOfAnApp(id string) error {
	url := fmt.Sprintf("/v1/apps/%s/relationships/betaTesters", id)
	headers := map[string]string{}
	_, status, err := a.client.Patch(url, headers, nil)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		return errors.New(strconv.Itoa(status))
	}
	return nil
}

// ReadTheBetaAppReviewDetailsResourceOfAnApp Get the beta app review details for a specific app.
func (a *Apps) ReadTheBetaAppReviewDetailsResourceOfAnApp(id string, parameters *QueryParameters) (*BetaAppReviewDetailWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/relationships/betaTesters", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response BetaAppReviewDetailWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadTheBetaLicenseAgreementOfAnApp Get the beta license agreement for a specific app.
func (a *Apps) ReadTheBetaLicenseAgreementOfAnApp(id string, parameters *QueryParameters) (*BetaLicenseAgreementWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/betaAppReviewDetail", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response BetaLicenseAgreementWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllBetaAppLocalizationsOfAnApp Get a list of localized beta test information for a specific app.
func (a *Apps) ListAllBetaAppLocalizationsOfAnApp(id string, parameters *QueryParameters) (*BetaLicenseAgreementWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/betaLicenseAgreement", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response BetaLicenseAgreementWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadTheXcodeCloudProductForAnApp Get the Xcode Cloud product information for an app you build with Xcode Cloud.
func (a *Apps) ReadTheXcodeCloudProductForAnApp(id string, parameters *QueryParameters) (*CiProductResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/ciProduct", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response CiProductResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllPricePointsForAnApp Get all the available price points for a specific app.
func (a *Apps) ListAllPricePointsForAnApp(id string, parameters *QueryParameters) (*AppPricePointsV3Response, error) {
	url := fmt.Sprintf("/v1/apps/%s/appPricePoints", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPricePointsV3Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadAppPricePointInformation Get details about a specific app price point.
func (a *Apps) ReadAppPricePointInformation(id string, parameters *QueryParameters) (*AppPricePointsV3Response, error) {
	url := fmt.Sprintf("/v3/appPricePoints/%s", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPricePointsV3Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAppPricePointEqualizations List all equivalent app prices points to a base price point.
func (a *Apps) ListAppPricePointEqualizations(id string, parameters *QueryParameters) (*AppPricePointsV3Response, error) {
	url := fmt.Sprintf("/v3/appPricePoints/%s/equalizations", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPricePointsV3Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllAppInfosForAnApp Get information about an app that is currently live on App Store, or that goes live with the next version.
func (a *Apps) ListAllAppInfosForAnApp(id string, parameters *QueryParameters) (*AppInfosResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appInfos", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppInfosResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllAppStoreVersionsForAnApp Get a list of all App Store versions of an app across all platforms.
func (a *Apps) ListAllAppStoreVersionsForAnApp(id string, parameters *QueryParameters) (*AppStoreVersionsResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appStoreVersions", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppStoreVersionsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadTheEndUserLicenseAgreementInformationOfAnApp Get the custom end user license agreement (EULA) for a specific app and the territories where the agreement applies.
func (a *Apps) ReadTheEndUserLicenseAgreementInformationOfAnApp(id string, parameters *QueryParameters) (*EndUserLicenseAgreementWithoutIncludesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/endUserLicenseAgreement", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response EndUserLicenseAgreementWithoutIncludesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllCustomProductPagesForAnApp Get a list of all custom product pages for a specific app.
func (a *Apps) ListAllCustomProductPagesForAnApp(id string, parameters *QueryParameters) (*AppCustomProductPagesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appCustomProductPages", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppCustomProductPagesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllAppStoreVersionExperimentsV2 ListAllAppStoreVersionExperimentsV2
func (a *Apps) ListAllAppStoreVersionExperimentsV2(id string, parameters *QueryParameters) (*AppStoreVersionExperimentsV2Response, error) {
	url := fmt.Sprintf("/v1/apps/%s/appStoreVersionExperimentsV2", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppStoreVersionExperimentsV2Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllPromotedPurchasesForAnApp Get a list of promoted in-app purchases, including promoted auto-renewable subscriptions, for an app.
func (a *Apps) ListAllPromotedPurchasesForAnApp(id string, parameters *QueryParameters) (*PromotedPurchasesResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/promotedPurchases", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response PromotedPurchasesResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetReviewSubmissionsForAnApp Get a list of review submissions associated with a specific app.
func (a *Apps) GetReviewSubmissionsForAnApp(id string, parameters *QueryParameters) (*ReviewSubmissionsResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/reviewSubmissions", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response ReviewSubmissionsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPowerAndPerformanceMetricsForAnApp Get the performance and power metrics data for the most recent version of an app.
func (a *Apps) GetPowerAndPerformanceMetricsForAnApp(id string, parameters *QueryParameters) (*XcodeMetrics, error) {
	url := fmt.Sprintf("/v1/apps/%s/perfPowerMetrics", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response XcodeMetrics
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllCustomerReviewsForAnApp Get a list of customer reviews for a specific app.
func (a *Apps) ListAllCustomerReviewsForAnApp(id string, parameters *QueryParameters) (*CustomerReviewsResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/customerReviews", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response CustomerReviewsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadPriceScheduleInformationForAnApp Read price schedule details for a specific app.
func (a *Apps) ReadPriceScheduleInformationForAnApp(id string, parameters *QueryParameters) (*AppPriceScheduleResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/appPriceSchedule", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPriceScheduleResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadAnAppsPriceScheduleInformation List the price schedule details for a specific app.
func (a *Apps) ReadAnAppsPriceScheduleInformation(id string, parameters *QueryParameters) (*AppPriceScheduleResponse, error) {
	url := fmt.Sprintf("/v1/appPriceSchedules/%s", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPriceScheduleResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAutomaticallyFeneratedPricesForAnApp List the automatically calculated prices for an app generated from a base territory.
func (a *Apps) ListAutomaticallyFeneratedPricesForAnApp(id string, parameters *QueryParameters) (*AppPricesV2Response, error) {
	url := fmt.Sprintf("/v1/appPriceSchedules/%s/automaticPrices", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPricesV2Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadTheBaseTerritoryForAnAppsPriceSchedule Read the base territory and currency for a specific app.
func (a *Apps) ReadTheBaseTerritoryForAnAppsPriceSchedule(id string, parameters *QueryParameters) (*TerritoryResponse, error) {
	url := fmt.Sprintf("/v1/appPriceSchedules/%s/baseTerritory", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response TerritoryResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListManuallyChosenPricesForAnApp List the prices you chose for a specific app.
func (a *Apps) ListManuallyChosenPricesForAnApp(id string, parameters *QueryParameters) (*AppPricesV2Response, error) {
	url := fmt.Sprintf("/v1/appPriceSchedules/%s/manualPrices", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPricesV2Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AddAcheduledPriceChangeToAnApp Create a scheduled price change for an app.
func (a *Apps) AddAcheduledPriceChangeToAnApp(parameters *AppPriceScheduleCreateRequest) (*AppPriceScheduleResponse, error) {
	url := fmt.Sprintf("/v1/appPriceSchedules")
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppPriceScheduleResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAvailabilityForAnApp Get a list of availabilities for a specific app.
func (a *Apps) ListAvailabilityForAnApp(id string, parameters *AppPriceScheduleCreateRequest) (*AppAvailabilityV2Response, error) {
	url := fmt.Sprintf("/v1/apps/%s/appAvailabilityV2", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AppAvailabilityV2Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadAnAppsAlternativeDistributionKey Get the alternative distribution keys for a specific app.
func (a *Apps) ReadAnAppsAlternativeDistributionKey(id string, parameters *QueryParameters) (*AlternativeDistributionKeyResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/alternativeDistributionKey", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response AlternativeDistributionKeyResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReadTheMarketplaceSearchDetailURL Get search detail URL for the alternative marketplace.
func (a *Apps) ReadTheMarketplaceSearchDetailURL(id string, parameters *QueryParameters) (*MarketplaceSearchDetailResponse, error) {
	url := fmt.Sprintf("/v1/apps/%s/alternativeDistributionKey", id)
	headers := map[string]string{}
	body, status, err := a.client.Patch(url, headers, parameters)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(strconv.Itoa(status))
	}
	var response MarketplaceSearchDetailResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
