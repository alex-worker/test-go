package ResourceManager

import (
	"os"
	"path/filepath"
)

type ResourceManager struct {
	resFolder string
}

func (r ResourceManager) GetResource(path string) (IResource, error) {
	filePath := filepath.Join(r.resFolder, path)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	res := FileResource{
		state: InternalResourceState{
			state:        Waiting,
			readyPercent: 0,
		},
		file: file,
	}
	return res, nil
}

func getResourceManager(dir string) IResourceManager {
	return ResourceManager{
		resFolder: dir,
	}
}
