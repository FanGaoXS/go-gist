package queue

type Queue interface {
	Size() int
	Empty() bool

	Push(val any) bool
	Pop() any
	Peek() any
}

type Type int

const (
	Array Type = iota
	Circular
)

func New(t Type, cap int) Queue {
	if t == Circular {
		if cap <= 0 {
			return nil
		}

		return &circular{
			elements: make([]any, cap), // 为数组赋值占位
			front:    0,
			rear:     0,
		}

	}
	return &array{
		elements: make([]any, 0),
		cap:      cap,
	}
}
