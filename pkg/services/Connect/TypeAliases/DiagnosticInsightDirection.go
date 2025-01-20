package TypeAliases

// DiagnosticInsight A string that describes the diagnostic insight direction.
type DiagnosticInsight string

const (
	DiagnosticInsightUp        DiagnosticInsight = "UP"        // The impact of this signature has regressed in the current version compared to previous versions.
	DiagnosticInsightDown      DiagnosticInsight = "DOWN"      // The impact of this signature has progressed in the current version compared to previous versions.
	DiagnosticInsightUndefined DiagnosticInsight = "UNDEFINED" // No significant change in impact of this signature in the current version compared to previous versions.
)

func (d DiagnosticInsight) String() string {
	return string(d)
}
