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

	v1 := e.Group("/v1")

	config.ConnectDb()

	route.AuthRoute(v1, config.DB)

	port := os.Getenv("PORT")

	e.Start(":" + port)
}
