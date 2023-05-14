package SDLInputSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/interfaces/IInputSystem"
)

func GetInputSystem() (IInputSystem, error) {
	keyboardState := sdl.GetKeyboardState()
	inputSystem := SDLInputSystem{
		keyboardState: keyboardState,
	}
	return &inputSystem, nil
}
