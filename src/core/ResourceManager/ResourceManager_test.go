package ResourceManager

import (
	"path/filepath"
	"runtime"
	"testing"
)
import "github.com/stretchr/testify/assert"

func getCurrentDir(t *testing.T) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return dir
}

func TestResourcesManagerForEmpty(t *testing.T) {
	myManager := getResourceManager("../../../data")
	res, err := myManager.GetResource("not existing path")
	assert.Equal(t, res, nil, "Must be equal")
	assert.NotEqualf(t, err, nil, "Must be equal")
}

func TestResourcesManager(t *testing.T) {
	myManager := getResourceManager("../../../data")

	res, err := myManager.GetResource("castletiles.tsx")
	assert.Equal(t, err, nil, "Must be equal")

	percent := res.GetReadyPercent()
	assert.Equal(t, percent, uint8(0), "Must be equal")

	state := res.GetState()
	assert.Equal(t, state, Waiting, "Must be equal")

	_, err = res.GetContent()
	assert.Equal(t, err.Error(), "not ready", "Must be equal")

	res.Load()

	percent = res.GetReadyPercent()
	assert.Equal(t, uint8(100), percent, "Must be equal")

	buf, err := res.GetContent()
	assert.Equal(t, err, nil, "Must be equal")
	assert.NotEqual(t, buf, nil, "Must be not equal")

	//t.Log(string(*buf)) for visual purpouse
}
