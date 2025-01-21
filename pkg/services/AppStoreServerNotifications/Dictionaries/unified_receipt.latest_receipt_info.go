package dictionaries

type UnifiedReceiptInfoLatest struct {
	// The AppAccountToken associated with this transaction.
	// This field is only present if your app supplied an
	// doc://com.apple.documentation/documentation/storekit/product/purchaseoption/3749440-appaccounttoken
	// or provided a UUID for the applicationUsername property when the user made the purchase.
	AppAccountToken string `json:"app_account_token"`
	// The time the App Store refunded a transaction or revoked it from family sharing,
	// in a date-time format similar to the ISO 8601.
	// This field is present only for refunded or revoked transactions.
	CancellationDate string `json:"cancellation_date"`
	// The time the App Store refunded a transaction or revoked it from family sharing,
	// in UNIX epoch time format, in milliseconds.
	// This field is present only for refunded or revoked transactions.
	// Use this time format for processing dates. For more information,
	// see cancellation_date_ms.
	CancellationDateMs string `json:"cancellation_date_ms"`
	// The time the App Store refunded a transaction or revoked it from family sharing, in Pacific Standard Time.
	// This field is present only for refunded or revoked transactions.
	CancellationDatePst string `json:"cancellation_date_pst"`
	// The reason for a refunded or revoked transaction.
	// A value of 1 indicates that the customer canceled their transaction due to an actual or perceived issue within your app.
	// A value of 0 indicates that the transaction was canceled for another reason,
	// for example, if the customer made the purchase accidentally.
	// Possible Values: 1, 0
	CancellationReason string `json:"cancellation_reason"`
	// The time a subscription expires or when it will renew, in a date-time format similar to the ISO 8601.
	ExpiresDate string `json:"expires_date"`
	// The time when a subscription expires or when it will renew, in UNIX epoch time format, in milliseconds.
	// Use this time format for processing dates. For more information,
	// see expires_date_ms.
	ExpiresDateMs string `json:"expires_date_ms"`
	// The time when a subscription expires or when it will renew, in Pacific Standard Time.
	ExpiresDatePst string `json:"expires_date_pst"`
	// A value that indicates whether the user is the purchaser of the product
	// or is a family member with access to the product through Family Sharing.
	// See in_app_ownership_type for more information.
	// Possible Values: FAMILY_SHARED, PURCHASED
	InAppOwnershipType string `json:"in_app_ownership_type"`
	// An indicator of whether an auto-renewable subscription is in the introductory price period. For more information, see is_in_intro_offer_period.
	// Possible Values: true, false
	IsInIntroOfferPeriod string `json:"is_in_intro_offer_period"`
	// An indicator of whether a subscription is in the free trial period.
	// For more information, see is_trial_period.
	// Possible Values: true, false
	IsTrialPeriod string `json:"is_trial_period"`
	// An indicator that the system canceled a subscription because the user upgraded.
	// This field is only present for subscription upgrade transactions.
	// Value: true
	IsUpgraded string `json:"is_upgraded"`
	// The reference name of a subscription offer you configured in App Store AppStoreConnectAPI.
	// This field is present when a customer redeemed a subscription offer code.
	// For more information, see offer_code_ref_name.
	OfferCodeRefName string `json:"offer_code_ref_name"`
	// The reference name of a subscription offer you configured in App Store AppStoreConnectAPI.
	// This field is present when a customer redeemed a subscription offer code.
	// For more information, see offer_code_ref_name.
	OriginalPurchaseDate string `json:"original_purchase_date"`
	// The time of the original app purchase, in UNIX epoch time format, in milliseconds.
	// Use this time format for processing dates.
	// This value indicates the date of the subscription’s initial purchase.
	// The original purchase date applies to all product types and remains the same in all transactions for the same product ID.
	// This value corresponds to the original transaction’s transactionDate property in StoreKit.
	OriginalPurchaseDateMs string `json:"original_purchase_date_ms"`
	// The time of the original app purchase, in Pacific Standard Time.
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	// The transaction identifier of the original purchase.
	// For more information, see original_transaction_id.
	OriginalTransactionId string `json:"original_transaction_id"`
	// The identifier of the subscription offer redeemed by the user. For more information,
	// see promotional_offer_id.
	PromotionalOfferId string `json:"promotional_offer_id"`
	// The unique identifier of the product purchased.
	// You provide this value when creating the product in App Store AppStoreConnectAPI,
	// and it corresponds to the productIdentifier property of the SKPayment
	// object stored in the transaction’s payment property.
	ProductId string `json:"product_id"`
	// The time when the App Store charged the user’s account for a subscription purchase or renewal after a lapse,
	// in a date-time format similar to the ISO 8601 standard.
	PurchaseDate string `json:"purchase_date"`
	// The time when the App Store charged the user’s account for a subscription purchase or renewal after a lapse,
	// in the UNIX epoch time format, in milliseconds. Use this time format for processing dates.
	PurchaseDateMs string `json:"purchase_date_ms"`
	// The time when the App Store charged the user’s account for a subscription purchase or renewal after a lapse,
	// in Pacific Standard Time.
	PurchaseDatePst string `json:"purchase_date_pst"`
	// The number of consumable products purchased.
	// This value corresponds to the quantity property of the SKPayment object stored in the transaction’s payment property.
	// The value is usually 1 unless modified with a mutable payment.
	// The maximum value is 10.
	Quantity string `json:"quantity"`
	// The identifier of the subscription group to which the subscription belongs.
	// The value for this field is identical to the subscriptionGroupIdentifier property in SKProduct.
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`
	// A unique identifier for a transaction such as a purchase, restore, or renewal.
	// For more information, see transaction_id.
	TransactionId string `json:"transaction_id"`
	// A unique identifier for purchase events across devices, including subscription-renewal events.
	// This value is the primary key to identify subscription purchases.
	WebOrderLineItemId string `json:"web_order_line_item_id"`
}
