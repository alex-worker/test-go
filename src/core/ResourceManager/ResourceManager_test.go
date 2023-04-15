package ResourceManager

import "testing"

func TestResourceManager(t *testing.T) {
	myManager := getResourceManager("/data")
	res := myManager.GetResource("castletiles.tsx")
	res.GetReadyPercent()
}
