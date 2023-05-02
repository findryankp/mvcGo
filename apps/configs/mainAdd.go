package configs

import "github.com/Findryankp/mvcGo/apps/utils"

func MainAdd() {
	utils.PackageIntall("github.com/joho/godotenv")
	lineNumber := utils.LineNumberGet("./main.go", `mvcGo.Init()`)
	utils.LineNumberInsertText("./main.go", `	mvcGo.Init()`+"\n"+MainContent(), lineNumber-1)
}

func MainContent() string {
	var text = `
	//run on port 8080
	LoadConfig(":8080")
`

	return text
}
