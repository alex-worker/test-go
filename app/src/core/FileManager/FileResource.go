package FileManager

import (
	"io"
	"os"
)

type FileResource struct {
	state  InternalResourceState
	file   *os.File
	buffer *[]byte
}

func (f *FileResource) Release() {
	f.closeFile()
	f.state.State = Closed
}

func (f *FileResource) GetContent() (*[]byte, error) {
	f.openFile()
	f.readAll()
	return f.buffer, nil
}

func (f *FileResource) closeFile() {
	if f.file != nil {
		err := f.file.Close()
		if err != nil {
			panic(err)
		}
		f.file = nil
	}
}

func (f *FileResource) openFile() {
	if f.file != nil {
		return
	}
	file, err := os.Open(f.state.FilePath)
	if err != nil {
		panic(err)
	}
	f.file = file
}

func (f *FileResource) readAll() {
	if f.file == nil {
		return
	}
	buffer, err := io.ReadAll(f.file)
	if err != nil {
		panic(err)
	}

	f.buffer = &buffer
	f.state.State = Ready
	f.state.ReadyPercent = 100
}
