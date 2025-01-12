package models

import "time"

type timestamp int64

func (t *timestamp) Time() time.Time {
	return time.UnixMilli(int64(*t))
}

// A token you use in a query to request the next set of transactions for the customer.
type revision string

// The start date of a timespan, expressed in UNIX time, in milliseconds.
type startDate timestamp

// The end date of a timespan, expressed in UNIX time, in milliseconds.
type endDate timestamp

// The unique identifier of the product, which you create in App Store Connect.
type productId string

// An optional filter that indicates the product type to include in the transaction history. Your query may specify more than one productType.
// Possible Values: AUTO_RENEWABLE, NON_RENEWABLE, CONSUMABLE, NON_CONSUMABLE
type productType string

// A string that describes whether the transaction was purchased by the customer, or is available to them through Family Sharing.
type inAppOwnershipType string

// An optional sort order for the transaction history records. The response sorts the transaction records by their recently modified date.
// The default value is ASCENDING, so you receive the oldest records first.
// Possible Values: ASCENDING, DESCENDING
type sort string

// An optional Boolean value that indicates whether the response includes only revoked transactions when the value is true,
// or contains only nonrevoked transactions when the value is false. By default, the request doesn’t include this parameter.
// Possible Values: true, false
type revoked bool

// An optional filter that indicates the subscription group identifier to include in the transaction history. Your query may specify more than one subscriptionGroupIdentifier.
type subscriptionGroupIdentifier string

// Set revoked to false to exclude revoked transactions instead.
// Possible Values: true, false
type excludeRevoked bool

// The unique identifier of an app in the App Store.
type appAppleId int64

// The bundle identifier of an app.
type bundleId string

// The server environment, either sandbox or production.
type environment string

// A Boolean value indicating whether the App Store has more transaction data.
type hasMore bool

// The original transaction identifier of a purchase.
type originalTransactionId string

// The status of an auto-renewable subscription.
// 1: The auto-renewable subscription is active.
// 2: The auto-renewable subscription is expired.
// 3: The auto-renewable subscription is in a billing retry period.
// 4: The auto-renewable subscription is in a Billing Grace Period.
// 5: The auto-renewable subscription is revoked. The App Store refunded the transaction or revoked it from Family Sharing.
type status int32

// The age of the customer’s account.
// 0: Account age is undeclared. Use this value to avoid providing information for this field.
// 1: Account age is between 0–3 days.
// 2: Account age is between 3–10 days.
// 3: Account age is between 10–30 days.
// 4: Account age is between 30–90 days.
// 5: Account age is between 90–180 days.
// 6: Account age is between 180–365 days.
// 7: Account age is over 365 days.
type accountTenure int32

// The UUID that an app optionally generates to map a customer’s In-App Purchase with its resulting App Store transaction.
type appAccountToken uuid

// A value that indicates the extent to which the customer consumed the in-app purchase.
// 0: The consumption status is undeclared. Use this value to avoid providing information for this field.
// 1: The in-app purchase is not consumed.
// 2: The in-app purchase is partially consumed.
// 3: The in-app purchase is fully consumed.
type consumptionStatus int32

// A Boolean value that indicates whether the customer consented to provide consumption data to the App Store.
type customerConsented bool

// A value that indicates whether the app successfully delivered an in-app purchase that works properly.
// 0: The app delivered the consumable in-app purchase and it’s working properly.
// 1: The app didn’t deliver the consumable in-app purchase due to a quality issue.
// 2: The app delivered the wrong item.
// 3: The app didn’t deliver the consumable in-app purchase due to a server outage.
// 4: The app didn’t deliver the consumable in-app purchase due to an in-game currency change.
// 5: The app didn’t deliver the consumable in-app purchase for other reasons.
type deliveryStatus int32

// A value that indicates the dollar amount of in-app purchases the customer has made in your app, since purchasing the app, across all platforms.
// 0: Lifetime purchase amount is undeclared. Use this value to avoid providing information for this field.
// 1: Lifetime purchase amount is 0 USD.
// 2: Lifetime purchase amount is between 0.01–49.99 USD.
// 3: Lifetime purchase amount is between 50–99.99 USD.
// 4: Lifetime purchase amount is between 100–499.99 USD.
// 5: Lifetime purchase amount is between 500–999.99 USD.
// 6: Lifetime purchase amount is between 1000–1999.99 USD.
// 7: Lifetime purchase amount is over 2000 USD.
type lifetimeDollarsPurchased int32

// A value that indicates the dollar amount of refunds the customer has received in your app, since purchasing the app, across all platforms.
// 0: Lifetime refund amount is undeclared. Use this value to avoid providing information for this field.
// 1: Lifetime refund amount is 0 USD.
// 2: Lifetime refund amount is between 0.01–49.99 USD.
// 3: Lifetime refund amount is between 50–99.99 USD.
// 4: Lifetime refund amount is between 100–499.99 USD.
// 5: Lifetime refund amount is between 500–999.99 USD.
// 6: Lifetime refund amount is between 1000–1999.99 USD.
// 7: Lifetime refund amount is over 2000 USD.
type lifetimeDollarsRefunded int32

// The platform on which the customer consumed the in-app purchase.
// 0: The platform is undeclared. Use this value to avoid providing information for this field.
// 1: An Apple platform.
// 2: Non-Apple platform.
type platform int32

// A value that indicates the amount of time that the customer used the app.
// 0: The engagement time is undeclared. Use this value to avoid providing information for this field.
// 1: The engagement time is between 0–5 minutes.
// 2: The engagement time is between 5–60 minutes.
// 3: The engagement time is between 1–6 hours.
// 4: The engagement time is between 6–24 hours.
// 5: The engagement time is between 1–4 days.
// 6: The engagement time is between 4–16 days.
// 7: The engagement time is over 16 days.
type playTime int32

// A value that indicates your preferred outcome for the refund request.
// 0: The refund preference is undeclared. Use this value to avoid providing information for this field.
// 1: You prefer that Apple grants the refund.
// 2: You prefer that Apple declines the refund.
// 3: You have no preference whether Apple grants or declines the refund.
type refundPreference int32

// A Boolean value that indicates whether you provided, prior to its purchase, a free sample or trial of the content,
// or information about its functionality.
type sampleContentProvided bool

// The status of a customer’s account within your app.
// 0: Account status is undeclared. Use this value to avoid providing information for this field.
// 1: The customer’s account is active.
// 2: The customer’s account is suspended.
// 3: The customer’s account is terminated.
// 4: The customer’s account has limited access.
type userStatus int32

type success string
