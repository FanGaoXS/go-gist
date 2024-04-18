package queue

type circular struct {
	elements []any
	front    int
	rear     int
}

func (c *circular) Size() int {
	return c.rear - c.front
}

func (c *circular) Empty() bool {
	return c.Size() == 0
}

func (c *circular) cap() int {
	return cap(c.elements)
}

func (c *circular) full() bool {
	return c.Size() == c.cap()
}

func (c *circular) Push(val any) bool {
	if c.full() {
		return false
	}

	i := c.rear % c.cap()
	c.elements[i] = val
	c.rear++
	return true
}

func (c *circular) Pop() any {
	if c.Empty() {
		return nil
	}

	i := c.front % c.cap()
	first := c.elements[i]
	c.front++
	return first
}

func (c *circular) Peek() any {
	if c.Empty() {
		return nil
	}

	i := c.front % c.cap()
	return c.elements[i]
}
