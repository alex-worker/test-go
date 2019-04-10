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

// Hero герои и прочие npc
type Hero struct {
	X uint32
	Y uint32
	Dir uint8
	Sprite Cell
}
