package SDLTimeSystem

import "github.com/veandco/go-sdl2/sdl"

type SDLTimeSystem struct {
	lastTime uint64
}

func (s *SDLTimeSystem) GetDeltaTime() uint64 {
	currentTime := sdl.GetTicks64()
	deltaTime := currentTime - s.lastTime
	s.lastTime = currentTime

	return deltaTime
}
