package configs

import "github.com/Findryankp/mvcGo/apps/utils"

func MainAdd() {
	utils.PackageIntall("github.com/joho/godotenv")
	lineNumber := utils.LineNumberGet("./main.go", `mvcGo.Init()`)
	utils.LineNumberInsertText("./main.go", MainContent(), lineNumber-1)
}

func MainContent() string {
	var text = `
	if err := mvcGo.Init(); err != nil {
		fmt.Println(err)
		return
	}

	//run on port 8080
	LoadConfig(":8080")
`

	return text
}
