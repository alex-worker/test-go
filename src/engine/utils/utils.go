package utils

import (
	"os"
	"path/filepath"
)

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
