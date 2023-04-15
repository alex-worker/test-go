package ResourceManager

type IResource interface {
	GetState() ResourceState
	GetReadyPercent() uint8
	GetContent() ([]byte, error)
	Load()
	Free()
}
