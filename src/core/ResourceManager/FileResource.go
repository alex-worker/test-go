package ResourceManager

import (
	"errors"
	"os"
)

type FileResource struct {
	state  InternalResourceState
	file   *os.File
	buffer []byte
}

func (f FileResource) GetState() ResourceState {
	return f.state.state
}

func (f FileResource) GetReadyPercent() uint8 {
	return f.state.readyPercent
}

func (f FileResource) Load() {
	f.buffer = os.
}

func (f FileResource) Free() {
	f.state.state = Closed
}

func (f FileResource) GetContent() ([]byte, error) {
	if f.state.state != Ready {
		return nil, errors.New("not ready")
	}
	return nil, nil
}
