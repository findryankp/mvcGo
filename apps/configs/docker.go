package configs

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/utils"
)

func DockerCreate() {
	file, err := utils.FilesCreate(".", "Dockerfile")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		utils.FilesAddContent(file, DockerContent())
		fmt.Println("setup docker success")
	}
}

func DockerContent() string {
	moduleName, _ := utils.ModuleNameGet()
	var text = `
	FROM golang:1.20-alpine

# membuat direktory app
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

EXPOSE 8080

# create executable
RUN go build -o ` + moduleName + `

# menjalankan file executablenya
CMD ["./` + moduleName + `"]
	`
	return text
}
