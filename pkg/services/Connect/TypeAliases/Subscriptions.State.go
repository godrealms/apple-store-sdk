package TypeAliases

type SubscriptionsStates string

const (
	SubscriptionsStateMissingMetadata          SubscriptionsStates = "MISSING_METADATA"
	SubscriptionsStateReadyToSubmit            SubscriptionsStates = "READY_TO_SUBMIT"
	SubscriptionsStateWaitingForReview         SubscriptionsStates = "WAITING_FOR_REVIEW"
	SubscriptionsStateInReview                 SubscriptionsStates = "IN_REVIEW"
	SubscriptionsStateDeveloperActionNeeded    SubscriptionsStates = "DEVELOPER_ACTION_NEEDED"
	SubscriptionsStatePendingBinaryApproval    SubscriptionsStates = "PENDING_BINARY_APPROVAL"
	SubscriptionsStateApproved                 SubscriptionsStates = "APPROVED"
	SubscriptionsStateDeveloperRemovedFromSale SubscriptionsStates = "DEVELOPER_REMOVED_FROM_SALE"
	SubscriptionsStateRemovedFromSale          SubscriptionsStates = "REMOVED_FROM_SALE"
	SubscriptionsStateRejected                 SubscriptionsStates = "REJECTED"
)

func (s SubscriptionsStates) String() string {
	return string(s)
}
