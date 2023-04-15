package ResourceManager

type IResourceManager interface {
	GetResource(path string) IResource
}
