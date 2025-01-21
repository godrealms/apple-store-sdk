package TypeAliases

// CiTestDestinationKind The string that represents the kind of a test destination.
type CiTestDestinationKind string

const (
	CiTestDestinationKindSimulator CiTestDestinationKind = "SIMULATOR" // The test destination is a simulated device.
	CiTestDestinationKindMac       CiTestDestinationKind = "MAC"       // The test destination is a Mac.
)

func (k CiTestDestinationKind) String() string {
	return string(k)
}
