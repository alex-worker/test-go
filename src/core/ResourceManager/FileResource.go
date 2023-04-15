package ResourceManager

import "errors"

type FileResource struct {
	state InternalResourceState
}

func (f FileResource) GetState() ResourceState {
	return f.state.state
}

func (f FileResource) GetReadyPercent() uint8 {
	return f.state.readyPercent
}

func (f FileResource) GetContent() ([]byte, error) {
	if f.state.state != Ready {
		return nil, errors.New("not ready")
	}
	return nil, nil
}

func (f FileResource) Free() {
	f.state.state = Closed
}
