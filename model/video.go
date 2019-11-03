package model

import "github.com/jinzhu/gorm"

type Video struct{
	gorm.Model
	Title string
	Info string `sql:"type:text"`
	URL string
}
