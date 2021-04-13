package models

import "gorm.io/gorm"

type Img struct {
	gorm.Model
	PostID int
	Url    string
}
