package queue

type array struct {
	elements []any
	cap      int
}

func (a *array) Size() int {
	return len(a.elements)
}

func (a *array) Empty() bool {
	return a.Size() == 0
}

// Push append to rear
func (a *array) Push(val any) bool {
	if a.cap > 0 && a.Size() == a.cap {
		return false
	}

	a.elements = append(a.elements, val)
	return true
}

// Pop from front
func (a *array) Pop() any {
	if a.Empty() {
		return nil
	}

	first := a.elements[0]
	a.elements = a.elements[1:]

	return first
}

func (a *array) Peek() any {
	if a.Empty() {
		return nil
	}

	return a.elements[0]
}
