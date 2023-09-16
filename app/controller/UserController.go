package controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rezaif79-ri/go-clean-architecture/intf"
)

type UserController struct {
	userService intf.UserService
}

func NewUserController(e *echo.Echo, userServiceObj intf.UserService) {
	userControllerObj := &UserController{
		userService: userServiceObj,
	}

	e.GET("printUserProfile", userControllerObj.printUserProfile)
	e.GET("testUserProfile", userControllerObj.testUserProfile)
}

func (uc *UserController) printUserProfile(ec echo.Context) error {
	fmt.Println("print user profile controller")
	return nil
}

func (uc *UserController) testUserProfile(ec echo.Context) error {
	fmt.Println("test user profile controller")
	uc.userService.TestUserProfile(ec.Request().Context())
	fmt.Println("== end of user profile controller")
	return nil
}
