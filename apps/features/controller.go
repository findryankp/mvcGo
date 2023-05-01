package features

import (
	"fmt"
	"os"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func ControllerCreate(controllerName string) {
	fileName := fmt.Sprintf("./controllers/%s.go", controllerName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, ControllerContent(controllerName))

	fmt.Println("Controller successfully created")
}

func ControllerContent(controllerName string) (text string) {
	moduleName, _ := utils.ModuleNameGet()
	text = `package controllers

import (
	"` + moduleName + `/configs"
	"` + moduleName + `/helpers"
	"` + moduleName + `/models"
	"github.com/labstack/echo/v4"
)

func ` + controllerName + `(c echo.Context) error {
	var data []models.` + controllerName + `
	if err := configs.DB.Find(&data); err != nil {
		return c.JSON(500, helpers.ResponseFail("err"))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func ` + controllerName + `ById(c echo.Context) error {
	var data models.` + controllerName + `
	if err := configs.DB.First(&data, c.Param("id")); err != nil {
		return c.JSON(500, helpers.ResponseFail("err"))
	}
	return c.JSON(200, helpers.ResponseSuccess("-", data))
}

func ` + controllerName + `Create(c echo.Context) error {
	var request models.` + controllerName + `Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModel` + controllerName + `(request)
	if err := configs.DB.Create(&data); err != nil {
		return c.JSON(500, helpers.ResponseFail("err"))
	}

	return c.JSON(201, helpers.ResponseSuccess("data created", nil))
}

func ` + controllerName + `Update(c echo.Context) error {
	var request models.` + controllerName + `Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(500, helpers.ResponseFail("-"))
	}

	data := models.RequestToModel` + controllerName + `(request)
	if err := configs.DB.Where("id = ?", c.Param("id")).Updates(&data); err != nil {
		return c.JSON(500, helpers.ResponseFail("err"))
	}

	return c.JSON(200, helpers.ResponseSuccess("data updated", nil))
}

func ` + controllerName + `Delete(c echo.Context) error {
	var data models.` + controllerName + `
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		return c.JSON(404, helpers.ResponseFail("data not found"))
	}

	configs.DB.Delete(&data)

	return c.JSON(200, helpers.ResponseSuccess("data deleted", nil))
}
	`

	return
}
