package math

import (
	"github.com/stretchr/testify/require"
	"test-go/src/defines"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("It working", func(t *testing.T) {
		s := defines.Size2D{
			Width:  10,
			Height: 10,
		}
		expected := New(s)
		require.Equal(t, expected.size, s, "size must be equal")

		mapSize := len(expected.data)
		require.Equal(t, defines.Dimension(mapSize), s.Width*s.Height, "size must be equal")
	})
}

func TestArray2D_GetCell(t *testing.T) {

}
