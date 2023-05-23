package math

import (
	"github.com/stretchr/testify/require"
	"test-go/src/defines"
	"testing"
)

func Test_Point2D(t *testing.T) {
	t.Run("It working", func(t *testing.T) {
		p := Point2D{
			X: 10,
			Y: 20,
		}
		require.Equal(t, p.X, defines.Dimension(10))
		require.Equal(t, p.Y, defines.Dimension(20))
	})
}
