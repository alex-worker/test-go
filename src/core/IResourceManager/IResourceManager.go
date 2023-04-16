package IResourceManager

type IResourceManager interface {
	GetResource(path string) (IResource, error)
	Release()
}
