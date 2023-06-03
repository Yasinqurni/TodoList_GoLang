package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	e := echo.New()

	port := os.Getenv("PORT")

	e.Start(":" + port)
}
