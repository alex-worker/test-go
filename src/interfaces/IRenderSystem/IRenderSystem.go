package IRenderSystem

type IRenderSystem interface {
	Draw()
	NewDrawable(DrawableType) IDrawable
}
