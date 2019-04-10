package def

// Cell ячейка карты
type Cell uint32

// Rect длина и ширина
type Rect struct {
	Width int
	Height int
}

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName string
	TileName string
	ScreenSize Rect
}
