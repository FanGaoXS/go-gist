package stack

type array struct {
	elements []any
}

func (s *array) Size() int {
	return len(s.elements)
}

func (s *array) Empty() bool {
	return s.Size() == 0
}

func (s *array) Push(o any) bool {
	s.elements = append(s.elements, o)
	return true
}

func (s *array) Pop() any {
	if s.Empty() {
		return nil
	}

	last := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]

	return last
}

func (s *array) Peek() any {
	if s.Empty() {
		return nil
	}

	return s.elements[s.Size()-1]
}
