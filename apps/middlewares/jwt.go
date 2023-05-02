package middlewares

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func JwtCreate() {
	utils.PackageIntall("github.com/golang-jwt/jwt")
	file, err := utils.FilesCreate("./helpers", "jwt.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(file, JwtContent())
		fmt.Println("setup jwt success")
	}
}

func JwtContent() string {
	var text = `package helpers

	import (
		"errors"
		"os"
		"strings"
		"time"
	
		"github.com/golang-jwt/jwt"
		"github.com/labstack/echo/v4"
	)
	
	var jwtKey = []byte(os.Getenv("SECRETKEY"))
	
	type JWTClaim struct {
		ID int ` + "`json:" + `"id"` + "`" + `
		jwt.StandardClaims
	}
	
	func GenerateToken(id int) (tokenString string, err error) {
		expirationTime := time.Now().Add(5 * time.Hour)
		claims := &JWTClaim{
			ID: id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = token.SignedString(jwtKey)
		return
	}
	
	func ValidateToken(signedToken string) (err error) {
		token, err := jwt.ParseWithClaims(
			signedToken,
			&JWTClaim{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtKey), nil
			},
		)
		if err != nil {
			return
		}
	
		claims, ok := token.Claims.(*JWTClaim)
		if !ok {
			err = errors.New("couldn't parse claims")
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			err = errors.New("token expired")
			return
		}
		return
	}
	
	func ClaimToken(c echo.Context) JWTClaim {
		secretKey, tokenString := HeaderToken(c)
		claims := &JWTClaim{}
		jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	
		return *claims
	}
	
	func HeaderToken(c echo.Context) (secretKey string, tokenString string) {
		secretKey = os.Getenv("SECRETKEY")
		reqToken := c.Request().Header.Get(os.Getenv("HEADERAUTH"))
		prefix := "Bearer "
		tokenString = strings.TrimPrefix(reqToken, prefix)
	
		return secretKey, tokenString
	}
	
	`
	return text
}
