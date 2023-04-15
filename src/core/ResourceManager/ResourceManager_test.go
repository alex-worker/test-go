package ResourceManager

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestResourceManager(t *testing.T) {
	myManager := getResourceManager("/data")
	res, err := myManager.GetResource("not existing path")

	assert.Equal(t, res, nil, "Must be equal")
	assert.Equal(t, err.Error(), "not found", "Must be equal")

	res, err = myManager.GetResource("castletiles.tsx")
	assert.Equal(t, err, nil, "Must be equal")

	percent := res.GetReadyPercent()
	assert.Equal(t, percent, uint8(0), "Must be equal")

	state := res.GetState()
	assert.Equal(t, state, Closed, "Must be equal")

	_, err = res.GetContent()
	assert.Equal(t, err.Error(), "not ready", "Must be equal")
}
