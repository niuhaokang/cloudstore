package database

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FileName string
	FileHash string
	FileAddr string
}

func (f *File) Insert(filename, filehash, fileAddr string) {}



