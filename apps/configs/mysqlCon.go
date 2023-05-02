package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func MysqlConCreate() {
	utils.PackageIntall("gorm.io/gorm")
	utils.PackageIntall("gorm.io/driver/mysql")

	fileConfig, err := utils.FilesCreate("./configs", "mysqlCon.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileConfig, MysqlConContent())
		fmt.Println("setup configs/mysqlCon success")
	}
}

func MysqlConContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `package configs

import (
	"fmt"
	"os"
	"` + moduleName + `/models"
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
	DB.AutoMigrate(
		&models.Users{},
	)
	fmt.Println("Migration done.")
}
	`

	return text
}
