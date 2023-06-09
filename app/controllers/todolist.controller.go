package controller

import "github.com/labstack/echo/v4"

type TodolistController interface {
	DeleteTodolist(c echo.Context) error
	GetAllTodolist(c echo.Context) error
	GetByIdTodolist(c echo.Context) error
}
