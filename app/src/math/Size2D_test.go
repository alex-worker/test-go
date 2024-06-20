package math

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"test-go/src/defines"
	"testing"
)

func getSize2D() Size2D {
	return Size2D{
		Width:  10,
		Height: 20,
	}
}

func TestSize2D_IsPointInto(t *testing.T) {
	t.Run("It working", func(t *testing.T) {
		p := getSize2D()
		point := Point2D{
			X: 5,
			Y: 5,
		}
		res := p.IsPointInto(point)
		assert.Equal(t, res, true)
		point = Point2D{
			X: 10,
			Y: 20,
		}
		res = p.IsPointInto(point)
		assert.Equal(t, res, false)
	})
}

func TestSize2D(t *testing.T) {
	t.Run("It working", func(t *testing.T) {
		p := getSize2D()
		require.Equal(t, p.Width, defines.Dimension(10))
		require.Equal(t, p.Height, defines.Dimension(20))
	})
}
