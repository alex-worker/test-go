package def

import (
	"path/filepath"
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

// GetPath возвращаем путь к файлу на основе имени файла и глобального пути
func GetPath(filename string) string {
	println("make path: ", filepath.Join(resFolder,filename) )
	return filepath.Join(resFolder,filename)
}
