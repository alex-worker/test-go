package defines

// Cell ячейка карты
type Cell uint32

// Layer слой
type Layer struct {
	Data *[][]Cell
	Name string
	W    uint32
	H    uint32
}

// Layers список слоев
type Layers []Layer

// Map прям вся карта ваще
type Map struct {
	Layers Layers
	W      uint32
	H      uint32
}

// AnimateFrame анимационный фрейм тайла
type AnimateFrame struct {
	Cell     Cell   // номер тайла
	Duration uint32 // задержка таймера
}

// AnimateTile анимированый тайл ну или нет...
type AnimateTile struct {
	Tick       int // текущий таймер
	Index      int // текущий индекс фрейма
	NeedUpdate bool
	Frames     []AnimateFrame // набор фреймов
}

// AnimateTiles набор тайлов
type AnimateTiles map[Cell]*AnimateTile
