package apps

func AddToMain() {
	GetPackage("github.com/joho/godotenv")
	lineNumber := GetLineNumber("./main.go", `mvcGo.Init()`)
	var text = `
if err := godotenv.Load(".env"); err != nil {
	log.Fatalf("Error loading .env file")
}

configs.InitDB()

e := echo.New()
routes.InitRouter(e)
e.Logger.Fatal(e.Start(":8080"))
`
	InsertAfterText("./main.go", `	mvcGo.Init()`+"\n"+text, lineNumber-1)
}
