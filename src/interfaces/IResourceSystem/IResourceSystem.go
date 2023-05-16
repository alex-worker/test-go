package IResourceSystem

type IResourceSystem interface {
	GetResource(path string) (IResource, error)
	Release()
}
