package route

import (
	controller "todoList_GoLang/app/controllers"
	repository "todoList_GoLang/app/repositories"
	service "todoList_GoLang/app/services"
	"todoList_GoLang/pkg"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRoute(e *echo.Echo, db *gorm.DB) {
	// Inisialisasi dependensi yang diperlukan
	userRepository := repository.NewUserRepositoryImpl(db)
	jwt := pkg.NewJWTService()
	authService := service.NewAuthServiceImpl(userRepository, jwt)
	userService := service.NewUserServiceImpl(userRepository)
	// Membuat instance controller dengan dependensi yang telah diinisialisasi
	authController := controller.NewAuthControllerImpl(authService, userService)

	// Rute register
	e.POST("/register", authController.Register)

	// Rute login
	e.POST("/login", authController.Login)

}
