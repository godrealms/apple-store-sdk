package TypeAliases

// CiExecutionProgress A string that represents the progress of an ongoing Xcode Cloud build.
type CiExecutionProgress string

const (
	CiExecutionProgressPending  CiExecutionProgress = "PENDING"  // Xcode Cloud hasn’t started the build.
	CiExecutionProgressRunning  CiExecutionProgress = "RUNNING"  // Xcode Cloud is performing the build.
	CiExecutionProgressComplete CiExecutionProgress = "COMPLETE" // Xcode Cloud completed the build.
)

func (e CiExecutionProgress) String() string {
	return string(e)
}
