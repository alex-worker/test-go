package ResourceManager

import "errors"

type FileResource struct {
	state        ResourceState
	readyPercent uint8
	path         string
}

func (f FileResource) GetState() ResourceState {
	return f.state
}

func (f FileResource) GetReadyPercent() uint8 {
	return f.readyPercent
}

func (f FileResource) GetContent() ([]byte, error) {
	if f.state != Ready {
		return nil, errors.New("nor ready")
	}
	return nil, nil
}

func (f FileResource) Free() {
	f.state = Closed
}
