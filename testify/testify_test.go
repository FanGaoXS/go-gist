package testify

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestify(t *testing.T) {
	s := require.New(t)

	var got = 4
	s.Equal(got, 4)
	//s.Equal <==>
	//if got != 4 {
	//	t.FailNow()
	//}
	//such as s.Equal(), s.Empty(), s.Len(), s.True() ...
}
