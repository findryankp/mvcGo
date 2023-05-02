package users

import (
	"fmt"
	"os"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func UsersControllerCreate() {
	fileName := fmt.Sprintf("./controllers/%s.go", "UsersController")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, UsersControllerConten())
}

func UsersControllerConten() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `package controllers

import (
	"` + moduleName + `/configs"
	"` + moduleName + `/helpers"
	"` + moduleName + `/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	query := configs.DB.Create(&user)
	if query.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", query.Error)
	}

	return helpers.ResponseJson(c, http.StatusCreated, true, "Registered succesfully", user)
}

func Login(c echo.Context) error {
	var user models.User
	var request models.User

	if err := c.Bind(&request); err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	record := configs.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
	}

	credentialError := helpers.CheckPassword(request.Password, user.Password)
	if !credentialError {
		return helpers.ResponseJson(c, http.StatusUnauthorized, false, "Invalid Credential", nil)
	}

	tokenString, err := helpers.GenerateToken(int(user.ID))
	if err != nil {
		return helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
	}

	response := map[string]string{
		"token": tokenString,
	}

	return helpers.ResponseJson(c, http.StatusOK, true, "-", response)
}
`

	return text
}
