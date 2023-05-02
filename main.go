package main

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func main() {
	filepath := "./configs/mysqlCon.go"
	lineNumber := utils.LineNumberGet(filepath, `AutoMigrate`)
	fmt.Println(lineNumber)
	utils.LineNumberInsertText("./configs/mysqlCon.go", `APAA KOE`, lineNumber)

}
