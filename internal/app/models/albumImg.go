package models

import "gorm.io/gorm"

type AlbumImg struct {
	gorm.Model
	AlbumID int
	Url    string
}

