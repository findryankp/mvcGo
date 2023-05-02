package users

import (
	"fmt"
	"os"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func UsersRouteCreate() {
	fileName := fmt.Sprintf("./routes/%s.go", "Users")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, UserModelContent())

	lineNumber := utils.LineNumberGet("./routes/initRoutes.go", `func InitRouter(e *echo.Echo)`)
	utils.LineNumberInsertText("./routes/initRoutes.go", `func InitRouter(e *echo.Echo){`+"\n"+`	`+"UsersRouter(e)", lineNumber-1)
}

func UserRouteContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `
package routes

import (
	"` + moduleName + `/controllers"

	"github.com/labstack/echo/v4"
)

func UsersRouter(e *echo.Echo) {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
}
	`
	return text
}
