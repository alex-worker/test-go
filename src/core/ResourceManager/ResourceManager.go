package ResourceManager

import (
	"errors"
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
		return nil, errors.New("not found")
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
	return ResourceManager{dir}
}
