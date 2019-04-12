package def

import (
	"path/filepath"
	"os"
	// "errors"
)

// Cell ячейка карты
type Cell uint32

// Tile анимированый тайл ну или нет...
type Tile struct {
	Tick uint32 // текущий таймер
	Tile Cell // текущий фрейм
	Frame []Cell // набор фреймов
}

// Pos координаты X Y
type Pos struct {
	X int
	Y int
}

// Rect длина и ширина
type Rect struct {
	Width int
	Height int
}

// LoadInfo структура хранения настроек игры
type LoadInfo struct {
	MapName string
	ResourceFolder string
	ScreenSize Rect
	FontName string
}

// ResourceFolder указатель на папку с ресурсами
var resFolder string

// SetResourceFolder устанавливаем путь к файлам
func SetResourceFolder(path string){
	resFolder = path
}

// GetPath для SDL-библиотек которые сами открывают файлы
func GetPath(filename string) string {
	return filepath.Join(resFolder,filename)
}

// OpenFile открываем файл
func OpenFile(filename string) (*os.File, error) {
	path := filepath.Join(resFolder,filename)
	return os.Open( path )
}