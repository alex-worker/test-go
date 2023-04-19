package SDLWindow

import (
	. "test-go/src/interfaces/IWindow"
)

type SDLWindow struct {
}

func (S *SDLWindow) Update() {
	//TODO implement me
	panic("implement me")
}

func GetWindow() (IWindow, error) {
	return &SDLWindow{}, nil
}
