package ResourceManager

import (
	"errors"
	"io"
	"os"
)

type FileResource struct {
	state  InternalResourceState
	file   *os.File
	buffer *[]byte
}

func (f *FileResource) GetState() ResourceState {
	return f.state.state
}

func (f *FileResource) GetReadyPercent() uint8 {
	return f.state.readyPercent
}

func (f *FileResource) Load() {
	var err error
	resState := []ResourceState{
		NotFound,
		Ready,
	}
	if StateIn(f.state.state, resState) {
		return
	}

	if f.file == nil {
		f.file, err = os.Open(f.state.filePath)
		if err != nil {
			panic(err)
		}
	}

	buffer, err := io.ReadAll(f.file)
	if err != nil {
		panic(err)
	}

	f.buffer = &buffer
	f.state.state = Ready
	f.state.readyPercent = 100
}

func (f *FileResource) Free() {
	if f.file != nil {
		err := f.file.Close()
		if err != nil {
			panic(err)
		}
		f.file = nil
	}
	f.state.state = Closed
}

func (f *FileResource) GetContent() (*[]byte, error) {
	if f.state.state != Ready {
		return nil, errors.New("not ready")
	}
	return f.buffer, nil
}
