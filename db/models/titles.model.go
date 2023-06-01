package model

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	UserId uint `validate:"required, user_id"`
	Title  string
}
