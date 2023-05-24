package def

import . "test-go/src/math"

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName        string
	ResourceFolder string
	ScreenSize     Size2D
	FontName       string
}
