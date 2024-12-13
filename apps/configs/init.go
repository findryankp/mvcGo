package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func InitConfig() {
	EnvironmentCreate()
	ResponseCreate()
	MysqlConCreate()
	MainAdd()
	PasswordCreate()
	LoadConfigCreate()
	TxtCreate()
}

func LoadConfigCreate() {
	fileConfig, err := utils.FilesCreate(".", "loadConfig.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileConfig, LoadConfigContent())
		fmt.Println("setup loadConfig success")
	}
}

func LoadConfigContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `package main

import (
	"log"

	"` + moduleName + `/configs"
	"` + moduleName + `/routes"
	"` + moduleName + `/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadConfig(port string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	configs.InitDB()

	e := echo.New()
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	routes.InitRouter(e)
	e.Logger.Fatal(e.Start(port))
}
	`
	return text
}
