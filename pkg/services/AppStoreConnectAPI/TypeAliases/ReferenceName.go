package TypeAliases

type ReferenceName string

func (receiver ReferenceName) String() string {
	return string(receiver)
}
