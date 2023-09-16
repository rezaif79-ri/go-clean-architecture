package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/go-clean-architecture/app/controller"
	"github.com/rezaif79-ri/go-clean-architecture/app/data"
	"github.com/rezaif79-ri/go-clean-architecture/app/service"
)

func main() {
	e := echo.New()

	// Init data layer
	userDL := data.NewUserDataLayer(nil)

	// Init services
	userService := service.NewUserService(userDL)

	// Init controller
	controller.NewUserController(e, userService)

	e.Logger.Info(e.Start("127.0.0.1:9090"))

}
