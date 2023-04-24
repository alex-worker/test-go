package TileAnimations

// AnimateTile анимированый тайл ну или нет...
type AnimateTile struct {
	Tick       int // текущий таймер
	Index      int // текущий индекс фрейма
	NeedUpdate bool
	Frames     []AnimateFrame // набор фреймов
}
