package users

import "github.com/Findryankp/mvcGo/apps/utils"

func UsersRouteCreate() {}

func UserRouteContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `
package routes

import (
	"` + moduleName + `/controllers"

	"github.com/labstack/echo/v4"
)

func authRouter(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/register", controllers.Register)
	g.POST("/login", controllers.Login)
}
	`

	return text
}
