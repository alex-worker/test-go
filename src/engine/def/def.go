package def

import (
	"os"
	"path/filepath"
	// "errors"
)

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

// Frame анимационный фрейм тайла
type Frame struct {
	Cell     Cell   // номер тайла
	Duration uint32 // задержка таймера
}

// Tile анимированый тайл ну или нет...
type Tile struct {
	Tick  uint32  // текущий таймер
	Index uint32  // текущий индекс фрейма
	Frame []Frame // набор фреймов
}

// Tiles набор тайлов
type Tiles map[int]Tile

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

// ResourceFolder указатель на папку с ресурсами
var resFolder string

// SetResourceFolder устанавливаем путь к файлам
func SetResourceFolder(path string) {
	resFolder = path
}

// GetPath для SDL-библиотек которые сами открывают файлы
func GetPath(filename string) string {
	return filepath.Join(resFolder, filename)
}

// OpenFile открываем файл
func OpenFile(filename string) (*os.File, error) {
	path := filepath.Join(resFolder, filename)
	return os.Open(path)
}
