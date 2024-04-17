package stack

type Stack interface {
	Size() int
	Empty() bool

	Push(val any) bool
	Pop() any
	Peek() any
}

func New() Stack {
	return &stackImpl{
		elements: make([]any, 0, 0),
	}
}

type stackImpl struct {
	elements []any
}

func (s *stackImpl) Size() int {
	return len(s.elements)
}

func (s *stackImpl) Empty() bool {
	return s.Size() == 0
}

func (s *stackImpl) Push(o any) bool {
	s.elements = append(s.elements, o)
	return true
}

func (s *stackImpl) Pop() any {
	if s.Empty() {
		return nil
	}

	last := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]

	return last
}

func (s *stackImpl) Peek() any {
	if s.Empty() {
		return nil
	}

	return s.elements[s.Size()-1]
}
