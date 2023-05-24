package def

import . "test-go/src/math"

// Pos координаты X Y
type Pos struct {
	X uint32
	Y uint32
}

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName        string
	ResourceFolder string
	ScreenSize     Size2D
	FontName       string
}
