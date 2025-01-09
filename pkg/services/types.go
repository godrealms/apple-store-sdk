package services

// NotificationType defines the type of Apple server notifications
type NotificationType string

// List of notification types
const (
	// NotificationTypeConsumptionRequest 一种通知类型，表明客户针对消费品应用内购买或自动续订订阅发起了退款请求，并且 App Store 要求您提供消费数据。
	//有关详细信息，请参阅发送消费信息。
	NotificationTypeConsumptionRequest NotificationType = "CONSUMPTION_REQUEST"

	// NotificationTypeDidChangeRenewalPref 一种通知类型及其子类型，表明客户对其订阅计划进行了更改。如果子类型是 UPGRADE，则用户升级了其订阅。
	//升级立即生效，开始新的计费周期，用户将收到上一周期未使用部分的按比例退款。如果子类型为 DOWNGRADE，则客户降级了其订阅。
	//降级将在下一个续订日期生效，并且不会影响当前有效的计划。如果子类型为空，则用户将其续订首选项更改回当前订阅，从而有效地取消降级。
	//有关订阅级别的更多信息，请参阅对组内的订阅进行排名。
	NotificationTypeDidChangeRenewalPref NotificationType = "DID_CHANGE_RENEWAL_PREF"

	// NotificationTypeDidChangeRenewalStatus 一种通知类型及其子类型，指示客户对订阅续订状态进行了更改。
	//如果子类型为 AUTO_RENEW_ENABLED，则客户重新启用订阅自动续订。如果子类型为 AUTO_RENEW_DISABLED，则客户禁用了订阅自动续订，
	//或者客户申请退款后 App Store 禁用了订阅自动续订。
	NotificationTypeDidChangeRenewalStatus NotificationType = "DID_CHANGE_RENEWAL_STATUS"

	// NotificationTypeDidFailToRenew 一种通知类型及其子类型，指示订阅由于计费问题而未能续订。订阅进入计费重试期。
	//如果子类型是 GRACE_PERIOD，则在宽限期内继续提供服务。如果子类型为空，则说明订阅不在宽限期内，您可以停止提供订阅服务。
	//通知客户他们的账单信息可能存在问题。 App Store 将在 60 天内继续重试计费，或者直到客户解决计费问题或取消订阅（以先到者为准）。
	//有关更多信息，请参阅减少非自愿订户流失。
	NotificationTypeDidFailToRenew NotificationType = "DID_FAIL_TO_RENEW"

	// NotificationTypeDidRenew 一种通知类型及其子类型，指示订阅已成功续订。如果子类型为BILLING_RECOVERY，则之前续订失败的过期订阅已成功续订。
	//如果子类型为空，则有效订阅已成功自动续订新的交易周期。为客户提供对订阅内容或服务的访问权限。
	NotificationTypeDidRenew NotificationType = "DID_RENEW"

	// NotificationTypeEXPIRED 一种通知类型及其子类型，指示订阅已过期。如果子类型为 VOLUNTARY，则订阅在用户禁用订阅续订后过期。
	//如果子类型为 BILLING_RETRY，则订阅已过期，因为计费重试期结束而没有成功的计费交易。
	//如果子类型为 PriceIncrease，则订阅已过期，因为客户不同意需要客户同意的价格上涨。
	//如果子类型为 PRODUCT_NOT_FOR_SALE，则订阅已过期，因为在订阅尝试续订时该产品不可购买。
	//没有子类型的通知表示订阅由于某些其他原因而过期。
	NotificationTypeEXPIRED NotificationType = "EXPIRED"

	// NotificationTypeExternalPurchaseToken 一种通知类型及其子类型 UNREPORTED，表明 Apple 为您的应用程序创建了外部购买令牌，但未收到报告。
	//有关报告令牌的更多信息，请参阅 externalPurchaseToken。此通知仅适用于使用外部购买来提供替代付款选项的应用程序。
	NotificationTypeExternalPurchaseToken NotificationType = "EXTERNAL_PURCHASE_TOKEN"

	// NotificationTypeGracePeriodExpired 一种通知类型，指示计费宽限期已结束而无需续订订阅，因此您可以关闭对服务或内容的访问。
	//通知客户他们的账单信息可能存在问题。 App Store 将在 60 天内继续重试计费，或者直到客户解决计费问题或取消订阅（以先到者为准）。
	//有关更多信息，请参阅减少非自愿订户流失。
	NotificationTypeGracePeriodExpired NotificationType = "GRACE_PERIOD_EXPIRED"

	// NotificationTypeOfferRedeemed 一种通知类型及其子类型，指示具有有效订阅的客户兑换了订阅优惠。
	//如果子类型为升级，则客户兑换了升级其有效订阅的优惠，该优惠立即生效。如果子类型为“降级”，则客户兑换了降级其有效订阅的优惠，该优惠将在下一个续订日期生效。
	//如果客户兑换了其有效订阅的优惠，您将收到不带子类型的 OfferRedeemed 通知类型。
	//有关订阅优惠代码的更多信息，请参阅在应用程序中支持订阅优惠代码。
	//有关促销优惠的更多信息，请参阅在您的应用程序中实施促销优惠。
	NotificationTypeOfferRedeemed NotificationType = "OFFER_REDEEMED"

	// NotificationTypeOneTimeCharge OneTimeCharge 通知目前仅在沙箱环境中可用。
	//一种通知类型，指示客户购买了消耗性、非消耗性或非续订订阅。
	//当客户通过家人共享获得对非消耗性产品的访问权限时，App Store 也会发送此通知。
	//有关自动续订订阅购买的通知，请参阅 SUBSCRIBED 通知类型。
	NotificationTypeOneTimeCharge NotificationType = "ONE_TIME_CHARGE"

	// NotificationTypePriceIncrease 一种通知类型及其子类型，表明系统已通知客户自动续订订阅价格上涨。
	//如果涨价需要客户同意，则如果客户未响应涨价，则子类型为“PENDING”；如果客户已同意涨价，则子类型为“ACCEPTED”。
	//如果涨价不需要客户同意，则子类型为已接受。
	//有关系统在显示需要客户同意的订阅价格上涨的价格同意表之前如何调用您的应用的信息，请参阅 paymentQueueShouldShowPriceConsent(_:)。
	//有关管理订阅价格的信息，请参阅管理自动续订订阅的价格上涨和管理价格。
	NotificationTypePriceIncrease NotificationType = "PRICE_INCREASE"

	// NotificationTypeRefund 一种通知类型，指示 App Store 已成功对消费品应用内购买、非消费品应用内购买、自动续订订阅或非续订订阅的交易进行退款。
	//reservationDate 包含退款交易的时间戳。 OriginalTransactionId 和 ProductId 标识原始交易和产品。 reservationReason 包含原因。
	//要请求客户的所有退款交易的列表，请参阅 App Store 服务器 API 中的获取退款历史记录。
	NotificationTypeRefund NotificationType = "REFUND"

	// NotificationTypeRefundDeclined 指示 App Store 拒绝退款请求的通知类型。
	NotificationTypeRefundDeclined NotificationType = "REFUND_DECLINED"

	// NotificationTypeRefundReversed 一种通知类型，表明 App Store 由于客户提出的争议而撤销了之前授予的退款。
	//如果您的应用因相关退款而撤销了内容或服务，则需要恢复它们。
	//此通知类型可适用于任何应用内购买类型：消耗型、非消耗型、非续订订阅和自动续订订阅。
	//对于自动续订订阅，当 App Store 撤销退款时，续订日期保持不变。
	NotificationTypeRefundReversed NotificationType = "REFUND_REVERSED"

	// NotificationTypeRenewalExtended 一种通知类型，指示 App Store 延长了特定订阅的订阅续订日期。
	//您可以通过调用 App Store Server API 中的“延长订阅续订日期”或“为所有活跃订阅者延长订阅续订日期”来请求订阅续订日期延期。
	NotificationTypeRenewalExtended NotificationType = "RENEWAL_EXTENDED"

	// NotificationTypeRenewalExtension 一种通知类型及其子类型，表明 App Store 正在尝试通过调用“为所有活跃订阅者延长订阅续订日期”来延长您请求的订阅续订日期。
	//如果子类型为“SUMMARY”，则 App Store 已完成为所有符合条件的订阅者延长续订日期。有关详细信息，请参阅responseBodyV2DecodedPayload 中的摘要。
	//如果子类型为 FAILURE，则特定订阅的续订日期延期未成功。详情请参见responseBodyV2DecodedPayload中的数据。
	NotificationTypeRenewalExtension NotificationType = "RENEWAL_EXTENSION"

	// NotificationTypeRevoke 指示客户有权通过“家人共享”进行应用内购买的通知类型不再可通过共享进行。
	//当购买者对其购买禁用“家庭共享”、购买者（或家庭成员）离开家庭群组或购买者收到退款时，App Store 会发送此通知。
	//您的应用程序还会收到 paymentQueue(_:didRevokeEntitlementsForProductIdentifiers:) 调用。
	//家庭共享适用于非消耗性应用内购买和自动续订订阅。有关家庭共享的更多信息，请参阅在应用程序中支持家庭共享。
	NotificationTypeRevoke NotificationType = "REVOKE"

	// NotificationTypeSubscribed 一种通知类型及其子类型，指示客户订阅了自动续订订阅。
	//如果子类型为 INITIAL_BUY，则客户首次通过家人共享购买或接收了对订阅的访问权限。
	//如果子类型是 RESUBSCRIBE，则用户通过家人共享重新订阅或接收对同一订阅或同一订阅组内的另一个订阅的访问权限。
	//有关其他产品类型购买的通知，请参阅 OneTimeCharge 通知类型。
	NotificationTypeSubscribed NotificationType = "SUBSCRIBED"

	// NotificationTypeTest 当您通过调用请求测试通知端点请求时，App Store 服务器发送的通知类型。
	//调用该端点来测试您的服务器是否正在接收通知。
	//仅当您提出请求时，您才会收到此通知。有关故障排除信息，请参阅获取测试通知状态端点。
	NotificationTypeTest NotificationType = "TEST"
)
