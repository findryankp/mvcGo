package features

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func MigrationCreate(modelName string) {
	filepath := "./configs/mysqlCon.go"
	lineNumber := utils.LineNumberGet(filepath, `DB.AutoMigrate(`)
	utils.LineNumberInsertText("./configs/mysqlCon.go", `	DB.AutoMigrate(`+"\n"+`		&models.`+modelName+"{},", lineNumber-1)
	fmt.Println("Migration successfully created")
}
