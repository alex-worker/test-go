package SDLInputSystem

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GetInputSystem() (*SDLInputSystem, error) {
	keyboardState := sdl.GetKeyboardState()
	inputSystem := SDLInputSystem{
		keyboardState: keyboardState,
	}
	return &inputSystem, nil
}
