package FileManager

import (
	"os"
	"path/filepath"
	"test-go/src/core/ResourceManager"
)

type FileManager struct {
	resFolder string
}

func (r FileManager) GetResource(path string) (ResourceManager.IResource, error) {
	filePath := filepath.Join(r.resFolder, path)
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	res := FileResource{
		state: ResourceManager.InternalResourceState{
			FilePath:     filePath,
			State:        ResourceManager.Waiting,
			ReadyPercent: 0,
		},
		file: nil,
	}
	return &res, nil
}

func getFileManager(dir string) ResourceManager.IResourceManager {
	return FileManager{
		resFolder: dir,
	}
}
