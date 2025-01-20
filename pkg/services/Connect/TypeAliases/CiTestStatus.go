package TypeAliases

// CiTestStatus A string that represents test status information.
type CiTestStatus string

const (
	CiTestStatusSuccess         CiTestStatus = "SUCCESS"          // The tests passed.
	CiTestStatusFailure         CiTestStatus = "FAILURE"          // The tests failed.
	CiTestStatusMixed           CiTestStatus = "MIXED"            // Some tests passed and some failed.
	CiTestStatusSkipped         CiTestStatus = "SKIPPED"          // Xcode Cloud skipped some tests.
	CiTestStatusExpectedFailure CiTestStatus = "EXPECTED_FAILURE" // Tests failed that you marked as expected to fail with [XCTExpectFailure](https://developer.apple.com/documentation/xctest/3726077- termxctexpectfailure).
)

func (s CiTestStatus) String() string {
	return string(s)
}
