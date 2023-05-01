package apps

import (
	"fmt"
	"os"
	"strings"
)

func CreateRoute(route string) {
	fileName := fmt.Sprintf("./routes/%sRouter.go", route)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	AddContentToFile(file, ContentRoute(route))

	lineNumber := GetLineNumber("./routes/initRoutes.go", `func InitRouter(e *echo.Echo)`)
	InsertAfterText("./routes/initRoutes.go", `func InitRouter(e *echo.Echo){`+"\n"+`	`+route+"Router(e)", lineNumber-1)

	fmt.Println("Routes successfully created")
}

func ContentRoute(routeName string) (text string) {
	moduleName, _ := GetModuleName()
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

func CreateInitRoute() string {
	GetPackage("github.com/labstack/echo/v4")
	moduleName, _ := GetModuleName()

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
