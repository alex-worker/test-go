package def

// Cell ячейка карты
type Cell uint32

// Rect длина и ширина
type Rect struct {
	Width uint32
	Height uint32
}

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName string
	TileName string
	ScreenSize Rect
}
