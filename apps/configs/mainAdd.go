package configs

import "github.com/Findryankp/mvcGo/apps/utils"

func MainAdd() {
	utils.PackageIntall("github.com/joho/godotenv")
	lineNumber := utils.LineNumberGet("./main.go", `mvcGo.Init()`)
	utils.LineNumberInsertText("./main.go", `	mvcGo.Init()`+"\n"+MainContent(), lineNumber-1)
}

func MainContent() string {
	var text = `
if err := godotenv.Load(".env"); err != nil {
	log.Fatalf("Error loading .env file")
}

configs.InitDB()

e := echo.New()
routes.InitRouter(e)
e.Logger.Fatal(e.Start(":8080"))
`

	return text
}
