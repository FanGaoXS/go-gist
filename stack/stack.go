package stack

type Type int

const (
	Array Type = iota
	Linked
)

type Stack interface {
	Size() int
	Empty() bool

	Push(val any) bool
	Pop() any
	Peek() any
}

func New(t Type) Stack {
	if t == Array {
		return &arrayStack{
			elements: make([]any, 0, 0),
		}
	}

	return &linkedStack{size: 0, top: nil}
}

type arrayStack struct {
	elements []any
}

func (s *arrayStack) Size() int {
	return len(s.elements)
}

func (s *arrayStack) Empty() bool {
	return s.Size() == 0
}

func (s *arrayStack) Push(o any) bool {
	s.elements = append(s.elements, o)
	return true
}

func (s *arrayStack) Pop() any {
	if s.Empty() {
		return nil
	}

	last := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]

	return last
}

func (s *arrayStack) Peek() any {
	if s.Empty() {
		return nil
	}

	return s.elements[s.Size()-1]
}

type linkedStackNode struct {
	val  any
	next *linkedStackNode
}

type linkedStack struct {
	size int
	top  *linkedStackNode
}

func (l *linkedStack) Size() int {
	return l.size
}

func (l *linkedStack) Empty() bool {
	return l.size == 0
}

func (l *linkedStack) Push(val any) bool {
	l.top = &linkedStackNode{
		val:  val,
		next: l.top,
	}
	l.size++

	return true
}

func (l *linkedStack) Pop() any {
	if l.Empty() {
		return nil
	}

	last := l.top
	l.top = l.top.next
	l.size--

	return last.val
}

func (l *linkedStack) Peek() any {
	if l.Empty() {
		return nil
	}

	return l.top.val
}
