package Engine

import (
	. "test-go/src/core/FileManager"
	"test-go/src/core/interfaces/IEngine"
	. "test-go/src/core/interfaces/IResourceManager"
)

type Engine struct {
	resourceManager IResourceManager
}

func (e *Engine) Run() {
}

func GetEngine() IEngine.IEngine {
	resourceManager := GetFileManager("/data")
	eng := &Engine{
		resourceManager: resourceManager,
	}
	return eng
}
