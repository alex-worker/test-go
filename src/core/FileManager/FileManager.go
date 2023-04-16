package FileManager

import (
	"os"
	"path/filepath"
	"test-go/src/core/IResourceManager"
)

type FileManager struct {
	resFolder string
}

func (r FileManager) GetResource(path string) (IResourceManager.IResource, error) {
	filePath := filepath.Join(r.resFolder, path)
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	res := FileResource{
		state: IResourceManager.InternalResourceState{
			FilePath:     filePath,
			State:        IResourceManager.Waiting,
			ReadyPercent: 0,
		},
		file: nil,
	}
	return &res, nil
}

func getFileManager(dir string) IResourceManager.IResourceManager {
	return FileManager{
		resFolder: dir,
	}
}
