package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	r := require.New(t)
	s := New(Array)

	r.Empty(s.Size())
	r.True(s.Empty())

	r.True(s.Push(1))
	r.True(s.Push(2))
	r.True(s.Push(3))
	r.True(s.Push(4))

	r.Equal(s.Size(), 4)
	r.Equal(s.Peek(), 4)
	r.Equal(s.Pop(), 4)
	r.Equal(s.Pop(), 3)
	r.Equal(s.Pop(), 2)

	r.Equal(s.Size(), 1)
	r.Equal(s.Peek(), 1)

	r.True(s.Push(100))
	r.True(s.Push(99))

	r.Equal(s.Size(), 3)
	r.Equal(s.Peek(), 99)

	r.Equal(s.Pop(), 99)
	r.Equal(s.Pop(), 100)
	r.Equal(s.Pop(), 1)

	r.Empty(s.Size())
	r.True(s.Empty())

	s = New(Linked)

	r.Empty(s.Size())
	r.True(s.Empty())

	r.True(s.Push(1))
	r.True(s.Push(2))
	r.True(s.Push(3))
	r.True(s.Push(4))

	r.Equal(s.Size(), 4)
	r.Equal(s.Peek(), 4)
	r.Equal(s.Pop(), 4)
	r.Equal(s.Pop(), 3)
	r.Equal(s.Pop(), 2)

	r.Equal(s.Size(), 1)
	r.Equal(s.Peek(), 1)

	r.True(s.Push(100))
	r.True(s.Push(99))

	r.Equal(s.Size(), 3)
	r.Equal(s.Peek(), 99)

	r.Equal(s.Pop(), 99)
	r.Equal(s.Pop(), 100)
	r.Equal(s.Pop(), 1)

	r.Empty(s.Size())
	r.True(s.Empty())
}
