package features

import (
	"github.com/Findryankp/mvcGo/apps/utils"
)

func AddMigration(modelName string) {
	filepath := "./configs/mysqlCon.go"
	lineNumber := utils.LineNumberGet(filepath, `DB.AutoMigrate(`)
	utils.LineNumberInsertText("./routes/initRoutes.go", `DB.AutoMigrate(`+"\n"+`	&models.`+modelName+"{}", lineNumber-1)
}
