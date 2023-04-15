package ResourceManager

type ResourceState uint8

const (
	Closed ResourceState = iota
	NotFound
	Waiting
	Loading
	Ready
)

type IResource interface {
	GetState() ResourceState
	GetReadyPercent() uint8
	GetContent() []byte
	Free()
}

type IResourceManager interface {
	GetResource(path string) IResource
}
