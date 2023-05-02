package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func PasswordCreate() {
	fileResponse, err := utils.FilesCreate("./helpers", "password.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(fileResponse, PasswordContent())
		fmt.Println("setup helpers/password success")
	}
}

func PasswordContent() string {
	var text = `package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "err"
	}

	password = string(bytes)
	return password
}

func CheckPassword(requestPassword, existPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existPassword), []byte(requestPassword))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
	`
	return text
}
