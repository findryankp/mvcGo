package users

import (
	"fmt"
	"os"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func UsersModelCreate() {
	fileName := fmt.Sprintf("./models/%s.go", "Users")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}

	utils.FilesAddContent(file, UsersModelContent())
}

func UsersModelContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `package models

import (
	"` + moduleName + `/helpers"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UsersRegister struct {
	Name     string ` + "`" + `json:"` + `name"` + ` form:"` + `name"` + "`" + `
	Email    string ` + "`" + `json:"` + `email"` + ` form:"` + `email"` + "`" + `
	Password string ` + "`" + `json:"` + `password"` + ` form:"` + `password"` + "`" + `
}

type UsersLogin struct {
	Email    string ` + "`" + `json:"` + `email"` + ` form:"` + `email"` + "`" + `
	Password string ` + "`" + `json:"` + `password"` + ` form:"` + `password"` + "`" + `
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password = helpers.EncryptPassword(user.Password)

	err = nil
	return
}
	`
	return text
}
