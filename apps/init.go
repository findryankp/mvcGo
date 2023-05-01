package apps

import "fmt"

func ExcecutionArgument(arg []string) {
	if arg[1] == "features" {
		if len(arg) != 3 {
			fmt.Println("wrong argument")
		} else {
			CreateModel(arg[2])
			CreateController(arg[2])
			CreateRoute(arg[2])
		}
	} else if arg[1] == "init" {
		CreateFolderIfNotExist()
		fmt.Println("setup folder success")

		fileEnv, err := CreateFile(".", ".env")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			AddContentToFile(fileEnv, CreateEnv())
			fmt.Println("setup env success")
		}

		fileConfig, err := CreateFile("./configs", "mysqlCon.go")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			AddContentToFile(fileConfig, CreateMysqlCon())
			fmt.Println("setup configs/mysqlCon success")
		}

		fileResponse, err := CreateFile("./helpers", "response.go")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			AddContentToFile(fileResponse, CreateResponse())
			fmt.Println("setup helpers/response success")
		}

		fileRoute, err := CreateFile("./routes", "initRoutes.go")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			AddContentToFile(fileRoute, CreateInitRoute())
			fmt.Println("setup routes/initRoutes success")
		}

		//add to Main function
		AddToMain()
	}
}
