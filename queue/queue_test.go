package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayQueue(t *testing.T) {
	r := require.New(t)
	q := New(Array, 4)

	r.Empty(q.Size())
	r.True(q.Empty())

	r.True(q.Push(1))
	r.True(q.Push(2))
	r.True(q.Push(3))
	r.True(q.Push(4))
	r.False(q.Push(5)) // is full

	r.Equal(q.Size(), 4)
	r.False(q.Empty())
	r.Equal(q.Peek(), 1)
	r.Equal(q.Pop(), 1)
	r.Equal(q.Pop(), 2)
	r.Equal(q.Pop(), 3)

	r.Equal(q.Size(), 1)
	r.Equal(q.Peek(), 4)

	r.True(q.Push(100))
	r.True(q.Push(99))

	r.Equal(q.Size(), 3)
	r.Equal(q.Peek(), 4)

	r.Equal(q.Pop(), 4)
	r.Equal(q.Pop(), 100)
	r.Equal(q.Pop(), 99)

	r.Empty(q.Size())
	r.True(q.Empty())

	q = New(Array, 0)

	r.Empty(q.Size())
	r.True(q.Empty())

	r.True(q.Push(1))
	r.True(q.Push(2))
	r.True(q.Push(3))
	r.True(q.Push(4))
	r.True(q.Push(5))

	r.Equal(q.Size(), 5)
	r.False(q.Empty())
	r.Equal(q.Peek(), 1)
	r.Equal(q.Pop(), 1)
	r.Equal(q.Pop(), 2)
	r.Equal(q.Pop(), 3)

	r.Equal(q.Size(), 2)
	r.Equal(q.Peek(), 4)

	r.True(q.Push(100))
	r.True(q.Push(99))

	r.Equal(q.Size(), 4)
	r.Equal(q.Peek(), 4)

	r.Equal(q.Pop(), 4)
	r.Equal(q.Pop(), 5)
	r.Equal(q.Pop(), 100)
	r.Equal(q.Pop(), 99)

	r.Empty(q.Size())
	r.True(q.Empty())
}

func TestCircularQueue(t *testing.T) {
	r := require.New(t)
	q := New(Circular, 1)

	r.Empty(q.Size())
	r.True(q.Empty())

	r.True(q.Push(1))
	r.False(q.Push(2)) // is full

	r.Equal(q.Size(), 1)
	r.False(q.Empty())
	r.Equal(q.Peek(), 1)
	r.Equal(q.Pop(), 1)

	r.Equal(q.Size(), 0)
	r.Nil(q.Peek())

	r.True(q.Push(100))

	r.Equal(q.Size(), 1)
	r.Equal(q.Peek(), 100)

	r.Equal(q.Pop(), 100)

	r.Empty(q.Size())
	r.True(q.Empty())

	q = New(Array, 4)

	r.Empty(q.Size())
	r.True(q.Empty())

	r.True(q.Push(1))
	r.True(q.Push(2))
	r.True(q.Push(3))
	r.True(q.Push(4))
	r.False(q.Push(5)) // is full

	r.Equal(q.Size(), 4)
	r.False(q.Empty())
	r.Equal(q.Peek(), 1)
	r.Equal(q.Pop(), 1)
	r.Equal(q.Pop(), 2)
	r.Equal(q.Pop(), 3)

	r.Equal(q.Size(), 1)
	r.Equal(q.Peek(), 4)

	r.True(q.Push(100))
	r.True(q.Push(99))

	r.Equal(q.Size(), 3)
	r.Equal(q.Peek(), 4)

	r.Equal(q.Pop(), 4)
	r.Equal(q.Pop(), 100)
	r.Equal(q.Pop(), 99)

	r.Empty(q.Size())
	r.True(q.Empty())
}
