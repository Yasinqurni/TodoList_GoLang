package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	UserId  uint
	TitleId uint
	List    string
	Done    bool
}
