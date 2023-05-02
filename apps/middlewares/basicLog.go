package middlewares

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func LogCreate() {
	utils.PackageIntall("github.com/labstack/echo/v4/middleware")
	file, err := utils.FilesCreate("./middlewares", "basicLog.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(file, LogContent())
		fmt.Println("setup log success")
	}
}

func LogContent() string {
	var text = `package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicLogger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
	`
	return text
}
