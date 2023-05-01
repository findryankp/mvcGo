package features

import (
	"fmt"
	"os"
	"strings"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func ModelCreate(modelName string) {
	fileName := fmt.Sprintf("./models/%s.go", modelName)

	// Attempt to open the file
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, ModelContent(modelName))

	fmt.Println("Model successfully created")
}

func ModelContent(modelName string) (text string) {
	lowerCase := strings.ToLower(modelName)
	text = `package models

import (
	"gorm.io/gorm"
)

type ` + modelName + ` struct {
	gorm.Model
	` + modelName + `Name string
}

type ` + modelName + `Request struct {
	` + modelName + `Name string ` + "`" + `json:"` + lowerCase + `_name" form:"` + lowerCase + `_name"` + "`" + `
}

type ` + modelName + `Response struct {
	Id       uint ` + "`" + `json:"id"` + "`" + `
	` + modelName + `Name string ` + "`" + `json:"` + lowerCase + `"` + "`" + `
}

func RequestToModel` + modelName + `(dt ` + modelName + `Request) ` + modelName + ` {
	return ` + modelName + `{
		` + modelName + `Name: dt.` + modelName + `Name,
	}
}

func ModelToResponse` + modelName + `(dt ` + modelName + `) ` + modelName + `Response {
	return ` + modelName + `Response{
		Id:           dt.ID,
		` + modelName + `Name: dt.` + modelName + `Name,
	}
}

func ListModelToResponse` + modelName + `(dt []` + modelName + `) []` + modelName + `Response {
	var responses = []` + modelName + `Response{}
	for _, v := range dt {
		responses = append(responses, ModelToResponse` + modelName + `(v))
	}
	return responses
}
`
	return
}
