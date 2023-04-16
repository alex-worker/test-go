package Engine

import (
	"test-go/src/core/FileManager"
	"test-go/src/core/IEngine"
	"test-go/src/core/IResourceManager"
)

type Engine struct {
	resourceManager IResourceManager.IResourceManager
}

func (e *Engine) Run() {
}

func GetEngine() IEngine.IEngine {
	resourceManager := FileManager.GetFileManager("/data")
	eng := &Engine{
		resourceManager: resourceManager,
	}
	return eng
}
