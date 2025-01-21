package TypeAliases

// DiagnosticInsightType A string that desribes the diagnostic insight type.
type DiagnosticInsightType string

const (
	DiagnosticInsightTypeTrend DiagnosticInsightType = "TREND" // Represents an insight type that indicates how the impact of signatures has changed between the current version and previous versions.
)

func (di DiagnosticInsightType) String() string {
	return string(di)
}
