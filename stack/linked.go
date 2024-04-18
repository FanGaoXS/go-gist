package stack

type linkedNode struct {
	val  any
	next *linkedNode
}

type linked struct {
	size int
	top  *linkedNode
}

func (l *linked) Size() int {
	return l.size
}

func (l *linked) Empty() bool {
	return l.size == 0
}

func (l *linked) Push(val any) bool {
	l.top = &linkedNode{
		val:  val,
		next: l.top,
	}
	l.size++

	return true
}

func (l *linked) Pop() any {
	if l.Empty() {
		return nil
	}

	last := l.top
	l.top = l.top.next
	l.size--

	return last.val
}

func (l *linked) Peek() any {
	if l.Empty() {
		return nil
	}

	return l.top.val
}
