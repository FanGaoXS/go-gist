package stack

type Stack interface {
	Size() int
	Empty() bool

	Push(val any) bool
	Pop() any
	Peek() any
}

type Type int

const (
	Array Type = iota
	Linked
)

func New(t Type) Stack {
	if t == Array {
		return &array{
			elements: make([]any, 0, 0),
		}
	}

	return &linked{
		size: 0,
		top:  nil,
	}
}
