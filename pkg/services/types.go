package services

// NotificationType defines the type of Apple server notifications
type NotificationType string

// List of notification types
const (
	NotificationTypeRenewal         NotificationType = "DID_RENEW"      // 订阅续订
	NotificationTypeCancel          NotificationType = "CANCEL"         // 订阅取消
	NotificationTypeInitialPurchase NotificationType = "INITIAL_BUY"    // 首次购买
	NotificationTypeRefund          NotificationType = "REFUND"         // 退款
	NotificationTypeGracePeriod     NotificationType = "GRACE_PERIOD"   // 宽限期
	NotificationTypePriceIncrease   NotificationType = "PRICE_INCREASE" // 价格变更
)

// NotificationPayload represents the payload sent by Apple in server-to-server notifications
type NotificationPayload struct {
	NotificationType       NotificationType `json:"notification_type"`         // 通知类型
	AutoRenewStatus        bool             `json:"auto_renew_status"`         // 是否自动续订
	AutoRenewProductID     string           `json:"auto_renew_product_id"`     // 自动续订的产品 ID
	Environment            string           `json:"environment"`               // 沙盒或生产环境
	BundleID               string           `json:"bid"`                       // 应用的 Bundle ID
	OriginalTransactionID  string           `json:"original_transaction_id"`   // 原始交易 ID
	ProductID              string           `json:"product_id"`                // 产品 ID
	ExpirationDate         string           `json:"expiration_date"`           // 订阅到期时间
	PurchaseDate           string           `json:"purchase_date"`             // 购买时间
	GracePeriodExpiresDate string           `json:"grace_period_expires_date"` // 宽限期到期时间
	PriceIncreaseStatus    int              `json:"price_increase_status"`     // 价格变更状态
}

// SubscriptionStatus represents the current status of a subscription
type SubscriptionStatus struct {
	IsActive       bool   `json:"is_active"`       // 当前订阅是否有效
	ExpirationDate string `json:"expiration_date"` // 到期时间
	ProductID      string `json:"product_id"`      // 产品 ID
	UserID         string `json:"user_id"`         // 用户 ID
}

// RefundDetails represents refund information
type RefundDetails struct {
	TransactionID string `json:"transaction_id"` // 退款相关的交易 ID
	Reason        string `json:"reason"`         // 退款原因
	RefundDate    string `json:"refund_date"`    // 退款日期
}
