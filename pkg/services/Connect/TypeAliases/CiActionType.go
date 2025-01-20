package TypeAliases

// CiActionType A string that represents the type of an Xcode Cloud workflow’s action.
type CiActionType string

const (
	CiActionTypeBuild   CiActionType = "BUILD"   // The action is a build action.
	CiActionTypeAnalyze CiActionType = "ANALYZE" // The action is an analyze action.
	CiActionTypeTest    CiActionType = "TEST"    // The action is a test action.
	CiActionTypeArchive CiActionType = "ARCHIVE" // The action is an archive action.
)

func (c CiActionType) String() string {
	return string(c)
}
