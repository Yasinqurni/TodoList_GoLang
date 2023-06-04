package controller

import (
	"net/http"
	"reflect"
	"strconv"
	service "todoList_GoLang/app/services"
	"todoList_GoLang/dto"
	"todoList_GoLang/pkg"
	helper "todoList_GoLang/response-helper"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type authControllerImpl struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthControllerImpl(authService service.AuthService, userService service.UserService) AuthController {
	return &authControllerImpl{authService: authService, userService: userService}
}

func (d *authControllerImpl) Register(c echo.Context) error {

	var req dto.UserRequest
	if err := c.Bind(&req); err != nil {
		data := helper.NewErrorResponse("cannot bind user request", err)
		return c.JSON(http.StatusBadRequest, data)
	}
	if req.Name == "" || reflect.TypeOf(req.Code).Kind() != reflect.Uint || len(strconv.Itoa(int(req.Code))) != 4 {
		data := helper.NewErrorResponse("please insert body", nil)
		return c.JSON(http.StatusBadRequest, data)
	}
	user, err := d.userService.GetByCode(strconv.Itoa(int(req.Code)))
	if err != nil {
		data := helper.NewErrorResponse("error search user", err)
		return c.JSON(http.StatusBadRequest, data)
	}
	if user != nil {
		data := helper.NewErrorResponse("duplicate user", err)
		return c.JSON(http.StatusBadRequest, data)
	}
	err = d.authService.Register(req.Name, strconv.Itoa(int(req.Code)))
	if err != nil {
		data := helper.NewErrorResponse("error register", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	data := helper.NewResponse("register successful", nil)
	return c.JSON(http.StatusOK, data)
}

func (d *authControllerImpl) Login(c echo.Context) error {

	var req dto.UserLogin

	if err := c.Bind(&req); err != nil {
		data := helper.NewErrorResponse("cannot bind user request", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	if reflect.TypeOf(req.Code).Kind() != reflect.Uint || len(strconv.Itoa(int(req.Code))) != 4 {
		data := helper.NewErrorResponse("please insert body", nil)
		return c.JSON(http.StatusBadRequest, data)
	}

	user, err := d.userService.GetByCode(strconv.Itoa(int(req.Code)))
	if err != nil {
		data := helper.NewErrorResponse("error search user", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	if user == nil {
		data := helper.NewErrorResponse("user not found", err)
		return c.JSON(http.StatusNotFound, data)
	}

	jwt := pkg.NewJWTService()

	token, err := jwt.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		data := helper.NewErrorResponse("error login", err)
		return c.JSON(http.StatusBadRequest, data)
	}

	data := helper.NewLoginResponse("register successful", token)
	return c.JSON(http.StatusOK, data)

}
