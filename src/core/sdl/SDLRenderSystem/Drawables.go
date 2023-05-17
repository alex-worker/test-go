package SDLRenderSystem

import (
	IRenderSystem "test-go/src/interfaces/IRenderSystem"
)

func (s *SDLRenderSystem) NewDrawable(drawableType IRenderSystem.DrawableType) IRenderSystem.IDrawable {
	switch drawableType {
	case IRenderSystem.ViewMap2D:
		return s.newViewMap2D()
	default:
		panic("unrecognized type")
	}

}

func (s *SDLRenderSystem) newViewMap2D() IRenderSystem.IViewMap2D {
	d := SDLViewMap2D{}
	return &d
}
