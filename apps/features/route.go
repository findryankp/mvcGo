package features

import (
	"fmt"
	"os"
	"strings"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func RouteCreate(route string) {
	fileName := fmt.Sprintf("./routes/%sRouter.go", route)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, RouteContent(route))

	lineNumber := utils.LineNumberGet("./routes/initRoutes.go", `func InitRouter(e *echo.Echo)`)
	utils.LineNumberInsertText("./routes/initRoutes.go", `func InitRouter(e *echo.Echo){`+"\n"+`	`+route+"Router(e)", lineNumber-1)

	fmt.Println("Routes successfully created")
}

func RouteContent(routeName string) (text string) {
	moduleName, _ := utils.ModuleNameGet()
	lowerCase := strings.ToLower(routeName)
	text = `package routes

import (
	"` + moduleName + `/controllers"
	"github.com/labstack/echo/v4"
)

func ` + routeName + `Router(e *echo.Echo) {
	g := e.Group("/` + lowerCase + `")
	g.GET("", controllers.` + routeName + `)
	g.GET("/:id", controllers.` + routeName + `ById)
	g.POST("", controllers.` + routeName + `Create)
	g.PUT("/:id", controllers.` + routeName + `Update)
	g.DELETE("/:id", controllers.` + routeName + `Delete)
}
	`
	return
}

func InitRouteCreate() {
	fileRoute, err := utils.FilesCreate("./routes", "initRoutes.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileRoute, InitRouteContent())
		fmt.Println("setup routes/initRoutes success")
	}
}

func InitRouteContent() string {
	utils.PackageIntall("github.com/labstack/echo/v4")
	moduleName, _ := utils.ModuleNameGet()

	var text = `package routes

import (
	"github.com/labstack/echo/v4"
	"` + moduleName + `/helpers"
)

func InitRouter(e *echo.Echo) {
	IndexRouter(e)
}

func IndexRouter(e *echo.Echo) {
	e.GET("/", index)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to mvcGo",
		"developmen_by": "Findryankp",
	}

	return c.JSON(200, helpers.ResponseSuccess("-", data))
}
	`
	return text
}
