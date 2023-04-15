package ResourceManager

type ResourceManager struct {
	resFolder string
}

func (r ResourceManager) GetResource(path string) IResource {
	res := FileResource{path: path}
	return res
}

func getResourceManager(dir string) IResourceManager {
	return ResourceManager{dir}
}
