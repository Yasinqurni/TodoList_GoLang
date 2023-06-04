package main

import (
	"os"
	route "todoList_GoLang/app/routers"
	config "todoList_GoLang/db/configs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	e := echo.New()

	config.ConnectDb()

	route.AuthRoute(e, config.DB)

	port := os.Getenv("PORT")

	e.Start(":" + port)
}
