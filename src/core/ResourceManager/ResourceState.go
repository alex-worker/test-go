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
	state        ResourceState
	readyPercent uint8
}
