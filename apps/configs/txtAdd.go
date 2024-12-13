package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func TxtCreate() {
	fileResponse, err := utils.FilesCreate(".", "model.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileResponse, PasswordContent())
		fmt.Println("Setup TXT Model success")
	}
}
