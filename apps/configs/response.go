package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func ResponseCreate() {
	fileResponse, err := utils.FilesCreate("./helpers", "response.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileResponse, ResponseContent())
		fmt.Println("setup helpers/response success")
	}
}

func ResponseContent() string {
	var text = `package helpers

func ResponseSuccess(message string, data any) map[string]any {
	return map[string]any{
		"status":  true,
		"message": message,
		"data":    data,
	}
}

func ResponseFail(message string) map[string]any {
	return map[string]any{
		"status":  false,
		"message": message,
	}
}
	`

	return text
}
