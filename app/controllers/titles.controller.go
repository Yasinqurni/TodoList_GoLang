package controller

import (
	"net/http"
	"strconv"
	service "todoList_GoLang/app/services"
	"todoList_GoLang/dto"
	helper "todoList_GoLang/response-helper"

	"github.com/labstack/echo/v4"
)

type TitleController interface {
	CreateTitle(c echo.Context) error
	UpdateTitle(c echo.Context) error
}

type titleControllerImpl struct {
	titleService service.TitleService
}

func NewTitleControllerImpl(titleService service.TitleService) TitleController {
	return &titleControllerImpl{titleService: titleService}
}

func (d *titleControllerImpl) CreateTitle(c echo.Context) error {

	var req dto.TitleRequest
	if err := c.Bind(&req); err != nil {
		data := helper.NewErrorResponse("cannot bind title request", err)
		return c.JSON(http.StatusBadRequest, data)
	}
	userId := c.Get("userId").(uint)

	err := d.titleService.Create(&req, userId)
	if err != nil {
		data := helper.NewErrorResponse("cannot create title", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	data := helper.NewResponse("create title successful", nil)
	return c.JSON(http.StatusOK, data)
}

func (d *titleControllerImpl) UpdateTitle(c echo.Context) error {

	var req dto.TitleRequest
	if err := c.Bind(&req); err != nil {
		data := helper.NewErrorResponse("cannot bind title request", err)
		return c.JSON(http.StatusBadRequest, data)
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := d.titleService.Update(req, uint(id))
	if err != nil {
		data := helper.NewErrorResponse("cannot update title", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	data := helper.NewResponse("create title successful", nil)
	return c.JSON(http.StatusOK, data)
}
