package SDLRenderSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/sdl"
	"test-go/src/defines"
	. "test-go/src/interfaces/IRenderSystem"
)

type SDLRenderSystem struct {
	renderer      *sdl.Renderer
	keyboardState []uint8
}

func (s *SDLRenderSystem) drawStart() {
	err := s.renderer.Clear()
	if err != nil {
		panic(err)
	}
}

func (s *SDLRenderSystem) drawEnd() {
	s.renderer.Present()
}

func (s *SDLRenderSystem) GetInput() defines.GameEvent {
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

func (s *SDLRenderSystem) Draw() {
	s.drawStart()
	s.drawEnd()
}

func GetRenderSystem(windowSize defines.Size) (IRenderSystem, error) {
	renderer := InitSDL(windowSize)
	keyboardState := sdl.GetKeyboardState()
	renderSystem := SDLRenderSystem{
		renderer:      renderer,
		keyboardState: keyboardState,
	}
	return &renderSystem, nil
}
