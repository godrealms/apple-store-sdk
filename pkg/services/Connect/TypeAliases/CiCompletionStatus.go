package TypeAliases

// CiCompletionStatus A string that represents the completion status of an Xcode Cloud build.
type CiCompletionStatus string

const (
	CiCompletionStatusSucceeded = "SUCCEEDED" // Xcode Cloud successfully completed a build.
	CiCompletionStatusFailed    = "FAILED"    // The Xcode Cloud build failed; for example, if you configure the Required to Pass setting for a test action and a unit test fails. For more information, see Add a Test Action in Configuring your Xcode Cloud workflow’s actions.
	CiCompletionStatusErrored   = "ERRORED"   // Xcode Cloud encountered an internal error when it performed the build.
	CiCompletionStatusCanceled  = "CANCELED"  // Xcode Cloud canceled the build because you manually canceled an ongoing build or because you enabled the Auto-cancel Builds setting for a workflow. For more information about the Auto-cancel Builds setting, see Xcode Cloud workflow reference.
	CiCompletionStatusSkipped   = "SKIPPED"   // Xcode Cloud skipped the build; for example, if you configure the Auto- termcancel Builds setting for a workflow and push many changes in quick succession.
)

func (s CiCompletionStatus) String() string {
	return string(s)
}
