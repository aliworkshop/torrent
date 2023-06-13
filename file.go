package torrent

import "github.com/anacrolix/torrent"

func NewFile(file *torrent.File) FileModel {
	return &File{File: file, name: file.DisplayPath(), path: file.Path(), length: file.Length()}
}

func (f *File) GetFile() *torrent.File {
	return f.File
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetPath() string {
	return f.path
}

func (f *File) GetLength() int64 {
	return f.length
}
