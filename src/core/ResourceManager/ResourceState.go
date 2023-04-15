package ResourceManager

type ResourceState uint8

const (
	Closed ResourceState = iota
	NotFound
	Waiting
	Loading
	Ready
)

type InternalResourceState struct {
	state        ResourceState
	readyPercent uint8
	path         string
}
