package FileManager

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	resFolder    string
	resourceList map[string]*FileResource
}

func (r *FileManager) Release() {
	for k, res := range r.resourceList {
		fmt.Println(k)
		res.Release()
	}
}

func (r *FileManager) GetResource(path string) (*FileResource, error) {
	filePath := filepath.Join(r.resFolder, path)
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	res := &FileResource{
		state: InternalResourceState{
			FilePath:     filePath,
			State:        Waiting,
			ReadyPercent: 0,
		},
		file: nil,
	}
	r.resourceList[filePath] = res
	return res, nil
}

func GetFileManager(dir string) (*FileManager, error) {
	return &FileManager{
		resFolder:    dir,
		resourceList: make(map[string]*FileResource),
	}, nil
}
