package apps

func CreateMysqlCon() string {

	GetPackage("gorm.io/gorm")
	GetPackage("gorm.io/driver/mysql")

	var text = `package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	USER := os.Getenv("DBUSERNAME")
	PASS := os.Getenv("DBPASS")
	HOST := os.Getenv("DBHOST")
	PORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate()
	fmt.Println("Migration done.")
}
	`

	return text
}

func AddMigration(modelName string) {
	lineNumber := GetLineNumber("./configs/mysqlCon.go", `fmt.Println("Migration done.")`)
	InsertAfterText("./configs/mysqlCon.go", `func InitRouter(e *echo.Echo){`+"\n"+`	`+modelName+"Router(e)", lineNumber-3)
}
