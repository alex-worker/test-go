package ResourceManager

import (
	"os"
)

type IResourceManager interface {
	GetPath(filename string) string
	OpenFile(filename string) (*os.File, error)
	CloseFile(file *os.File)
}
