package IResourceManager

type IResource interface {
	GetContent() (*[]byte, error)
	Release()
}
