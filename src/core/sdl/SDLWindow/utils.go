package SDLWindow

import (
	"github.com/veandco/go-sdl2/sdl"
)

func drawStart(renderer *sdl.Renderer) {
	err := renderer.Clear()
	if err != nil {
		panic(err)
	}
}

func drawEnd(renderer *sdl.Renderer) {

	//if isShowFps {
	//	showFPS()
	//}
	//
	renderer.Present()

	//currentTime := sdl.GetTicks64()
	//deltaTime := currentTime - startTicks

	//fmt.Println(deltaTime)

	//var fps uint64

	//if deltaTime > 0 {
	//	fps = 1000 / deltaTime
	//} else {
	//	fps = 0
	//}
	//
	//fmt.Println("fps: ", fps)
	//
	//lastTime = currentTime
	//
	//UpdateUI(deltaTime)
}
