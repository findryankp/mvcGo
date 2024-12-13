package features

import (
	"fmt"
	"os"
	"strings"

	"github.com/Findryankp/mvcGo/apps/utils"
	"github.com/Findryankp/mvcGo/helpers"
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

func ModelCreateFromTxt(modelName string) {
	fileName := fmt.Sprintf("./models/%s.go", modelName)

	// Attempt to open the file
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, ModelContentFromTxt(modelName))

	fmt.Println("Model successfully created")
}

func ModelContentFromTxt(modelName string) (text string) {
	// lowerCase := strings.ToLower(modelName)
	getStr := helpers.GetStringFromTxt()
	fields := helpers.ExplodeStringFromTxt(getStr)

	var structFields, requestFields, responseFields, requestToModelFields, modelToResponseFields []string

	for _, field := range fields {
		name := strings.Title(field[0])
		typeDef := field[1]
		jsonTag := helpers.CamelToSnake(field[0])

		// Add to struct fields
		structFields = append(structFields, "\t"+name+" "+typeDef)

		// Add to request struct fields
		requestFields = append(requestFields, "\t"+name+" "+typeDef+" `json:\""+jsonTag+"\" form:\""+jsonTag+"\"`")

		// Add to response struct fields
		responseFields = append(responseFields, "\t"+name+" "+typeDef+" `json:\""+jsonTag+"\"`")

		requestToModelFields = append(requestToModelFields, "\t\t"+name+": dt."+name+",")
		modelToResponseFields = append(modelToResponseFields, "\t\t"+name+": dt."+name+",")
	}

	text = `package models

import (
	"gorm.io/gorm"
)

type ` + modelName + ` struct {
	gorm.Model
` + strings.Join(structFields, "\n") + `
}

type ` + modelName + `Request struct {
` + strings.Join(requestFields, "\n") + `
}

type ` + modelName + `Response struct {
	Id uint ` + "`json:\"id\"`" + `
` + strings.Join(responseFields, "\n") + `
}

func RequestToModel` + modelName + `(dt ` + modelName + `Request) ` + modelName + ` {
	return ` + modelName + `{
` + strings.Join(requestToModelFields, "\n") + `
	}
}

func ModelToResponse` + modelName + `(dt ` + modelName + `) ` + modelName + `Response {
	return ` + modelName + `Response{
		Id: dt.ID,
` + strings.Join(modelToResponseFields, "\n") + `
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

	return text
}
