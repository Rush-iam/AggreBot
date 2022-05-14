package user_state

type UserState struct {
	State int
	Value int64
}

const (
	Empty = iota
	SourceAdd
	SourceRename
	FilterSet
)
