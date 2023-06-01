package config

import (
	"time"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/pkg"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	database := pkg.GormInit()
	err := database.AutoMigrate(
		&model.User{},
		&model.Activity{},
		&model.Title{},
	)
	if err != nil {
		pkg.Logger(err.Error())
	}
	dbPool, err := database.DB()
	if err != nil {
		pkg.Logger(err.Error())
	}
	dbPool.SetMaxIdleConns(3)
	dbPool.SetMaxOpenConns(30)
	dbPool.SetConnMaxLifetime(30 * time.Minute)
	DB = database
}
