package model

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	UserId     uint
	Title      string
	Activities []Activity
}
