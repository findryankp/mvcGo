package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func EnvironmentCreate() {
	fileEnv, err := utils.FilesCreate(".", ".env")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileEnv, EnvironmentContent())
		fmt.Println("setup env success")
	}
}

func EnvironmentContent() string {
	var text = `DBUSERNAME=findryankp
DBPASS=findryankp
DBHOST=localhost
DBPORT=3306
DBNAME=testdb
	`
	return text
}
