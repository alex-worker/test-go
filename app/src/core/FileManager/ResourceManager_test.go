package FileManager

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
	myManager, err := GetFileManager("../../../data")
	assert.NoError(t, err, "")
	res, err := myManager.GetResource("not existing path")
	assert.Equal(t, res, nil, "Must be equal")
	assert.NotEqualf(t, err, nil, "Must be equal")

	myManager.Release()
}

func TestResourcesManager(t *testing.T) {
	myManager, err := GetFileManager("../../../data")
	assert.NoError(t, err, "")

	res, err := myManager.GetResource("castletiles.tsx")
	assert.Equal(t, err, nil, "Must be equal")

	buf, err := res.GetContent()
	assert.Equal(t, err, nil, "Must be equal")
	assert.NotEqual(t, buf, nil, "Must be not equal")

	//t.Log(string(*buf)) for visual purpouse
	myManager.Release()
}
