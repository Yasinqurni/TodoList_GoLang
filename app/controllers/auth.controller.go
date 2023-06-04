package controller

import (
	"net/http"
	"reflect"
	"strconv"
	service "todoList_GoLang/app/services"
	"todoList_GoLang/dto"
	"todoList_GoLang/pkg"

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
		return c.JSON(http.StatusBadRequest, err)
	}
	if req.Name == "" || reflect.TypeOf(req.Code).Kind() != reflect.Uint || len(strconv.Itoa(int(req.Code))) != 4 {
		return c.JSON(http.StatusBadRequest, "error body")
	}
	user, err := d.userService.GetByCode(strconv.Itoa(int(req.Code)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if user != nil {
		return c.JSON(http.StatusBadRequest, "duplicate user")
	}
	err = d.authService.Register(req.Name, strconv.Itoa(int(req.Code)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (d *authControllerImpl) Login(c echo.Context) error {

	var req dto.UserLogin

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if reflect.TypeOf(req.Code).Kind() != reflect.Uint || len(strconv.Itoa(int(req.Code))) != 4 {
		return c.JSON(http.StatusBadRequest, "error body")
	}

	user, err := d.userService.GetByCode(strconv.Itoa(int(req.Code)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	jwt := pkg.NewJWTService()

	token, err := jwt.GenerateToken(strconv.Itoa(int(user.ID)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, token)

}
