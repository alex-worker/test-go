package SDLRenderSystem

import (
	. "test-go/src/interfaces/IRenderSystem"
)

func (s *SDLRenderSystem) NewDrawable(drawableType DrawableType) IDrawable {
	switch drawableType {
	case ViewMap2D:
		return s.newViewMap2D()
	default:
		panic("unsupported type")
	}
}

func (s *SDLRenderSystem) newViewMap2D() IViewMap2D {
	d := SDLViewMap2D{}
	return &d
}
