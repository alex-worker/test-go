package def

// Pos координаты X Y
type Pos struct {
	X uint32
	Y uint32
}

// Size длина и ширина
type Size struct {
	Width  uint32
	Height uint32
}

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName        string
	ResourceFolder string
	ScreenSize     Size
	FontName       string
}
