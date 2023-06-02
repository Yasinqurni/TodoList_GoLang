package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	UserId  uint `validate:"required, user_id"`
	TitleId uint `validate:"required, title_id"`
	List    string
	Done    bool
	Titles  []Title
}
