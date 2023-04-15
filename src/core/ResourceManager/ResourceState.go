package ResourceManager

type ResourceState uint8

const (
	Closed ResourceState = iota
	NotFound
	Waiting // found
	Loading // loading some percent
	Ready   // full loading and ready
)

type InternalResourceState struct {
	filePath     string
	state        ResourceState
	readyPercent uint8
}

func StateIn(s ResourceState, arr []ResourceState) bool {
	for _, item := range arr {
		if item == s {
			return true
		}
	}
	return false
}
