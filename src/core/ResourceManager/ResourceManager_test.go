package ResourceManager

import "testing"

func TestResourceManager(t *testing.T) {
	myManager := getResourceManager("/data")
	myManager.GetResource("castletiles.tsx")
}
