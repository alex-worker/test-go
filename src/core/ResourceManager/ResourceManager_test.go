package ResourceManager

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestResourceManager(t *testing.T) {
	myManager := getResourceManager("/data")
	res := myManager.GetResource("castletiles.tsx")
	percent := res.GetReadyPercent()
	assert.Equal(t, percent, uint8(0), "Must be equal")
	state := res.GetState()
	assert.Equal(t, state, Closed, "Must be equal")
	_, err := res.GetContent()
	assert.Equal(t, err.Error(), "not ready", "Must be equal")
}
