package ResourceManager

import (
	"os"
	"path/filepath"
)

type ResourceManager struct {
	resFolder string
}

func (r ResourceManager) GetPath(filename string) string {
	return filepath.Join(r.resFolder, filename)
}

func (r ResourceManager) OpenFile(filename string) (*os.File, error) {
	path := filepath.Join(r.resFolder, filename)
	return os.Open(path)
}

func (r ResourceManager) CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
