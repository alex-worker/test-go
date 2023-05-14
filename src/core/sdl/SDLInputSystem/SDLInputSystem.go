package SDLInputSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	"test-go/src/defines"
	. "test-go/src/interfaces/IInputSystem"
)

type SDLInputSystem struct {
	keyboardState []uint8
}

func (s *SDLInputSystem) GetInput() defines.GameEvent {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		default:
			//fmt.Println(event)
			break
		case *sdl.QuitEvent:
			return defines.EventQuit
		}
	}

	if s.keyboardState[sdl.SCANCODE_UP] != 0 {
		return defines.EventPressUp
	} else if s.keyboardState[sdl.SCANCODE_DOWN] != 0 {
		return defines.EventPressDown
	} else if s.keyboardState[sdl.SCANCODE_LEFT] != 0 {
		return defines.EventPressLeft
	} else if s.keyboardState[sdl.SCANCODE_RIGHT] != 0 {
		return defines.EventPressRight
	}

	return defines.EventNo
}

func GetInputSystem() (IInputSystem, error) {
	keyboardState := sdl.GetKeyboardState()
	inputSystem := SDLInputSystem{
		keyboardState: keyboardState,
	}
	return &inputSystem, nil
}
