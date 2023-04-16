package FileManager

import (
	"fmt"
	"os"
	"path/filepath"
	"test-go/src/core/IResourceManager"
)

type FileManager struct {
	resFolder    string
	resourceList map[string]IResourceManager.IResource
}

func (r *FileManager) Release() {
	for k, res := range r.resourceList {
		fmt.Println(k)
		res.Release()
	}
}

func (r *FileManager) GetResource(path string) (IResourceManager.IResource, error) {
	filePath := filepath.Join(r.resFolder, path)
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	res := &FileResource{
		state: IResourceManager.InternalResourceState{
			FilePath:     filePath,
			State:        IResourceManager.Waiting,
			ReadyPercent: 0,
		},
		file: nil,
	}
	r.resourceList[filePath] = res
	return res, nil
}

func getFileManager(dir string) IResourceManager.IResourceManager {
	return &FileManager{
		resFolder:    dir,
		resourceList: make(map[string]IResourceManager.IResource),
	}
}
