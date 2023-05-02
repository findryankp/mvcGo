package middlewares

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func AuthenticationCreate() {
	file, err := utils.FilesCreate("./middlewares", "authentication.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(file, AuthenticationContent())
		fmt.Println("setup authentication success")
	}
}

func AuthenticationContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `package middlewares

	import (
		"net/http"
		"os"
		"` + moduleName + `/helpers"
		"github.com/labstack/echo/v4"
	)
	
	func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, tokenString := helpers.HeaderToken(c)
	
			if token := c.Request().Header.Get(os.Getenv("HEADERAUTH")); tokenString == "" || tokenString == token {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("request does not contain an access token"))
			}
	
			if err := helpers.ValidateToken(tokenString); err != nil {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("request does not contain an valid token"))
			}
			return next(c)
		}
	}
	
	`
	return text
}
