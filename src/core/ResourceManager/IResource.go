package ResourceManager

type IResource interface {
	GetContent() (*[]byte, error)
	Release()
}
