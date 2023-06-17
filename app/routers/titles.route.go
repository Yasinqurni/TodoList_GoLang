package route

import (
	repository "todoList_GoLang/app/repositories"
)

func TitleRoute(e *echo.Group, db *gorm.DB) {
	titleRepository := repository.NewTitleRepositoryImpl(db)
}