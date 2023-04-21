package SDLWindow

import (
	"github.com/veandco/go-sdl2/sdl"
	"test-go/src/defines"
)

type SDLWindow struct {
	size          defines.Size
	renderer      *sdl.Renderer
	keyboardState []uint8
}

func (s *SDLWindow) DrawStart() {
	drawStart(s.renderer)
}

func (s *SDLWindow) DrawEnd() {
	drawEnd(s.renderer)
}

func (s *SDLWindow) GetInput() defines.GameEvent {
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

func GetWindow(size defines.Size) (*SDLWindow, error) {
	renderer := initSDL(size)
	keyboardState := sdl.GetKeyboardState()
	return &SDLWindow{
		size:          size,
		renderer:      renderer,
		keyboardState: keyboardState,
	}, nil
}
