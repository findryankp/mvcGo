package middlewares

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func CorsCreate() {
	file, err := utils.FilesCreate("./middlewares", "cors.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(file, CorsContent())
		fmt.Println("setup authentication success")
	}
}

func CorsContent() string {
	var text = `package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Cors(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}
	`

	return text
}
