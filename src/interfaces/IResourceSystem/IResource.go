package IResourceSystem

type IResource interface {
	GetContent() (*[]byte, error)
	Release()
}
