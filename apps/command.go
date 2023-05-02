package apps

import (
	"fmt"

	"github.com/Findryankp/mvcGo/apps/configs"
	"github.com/Findryankp/mvcGo/apps/features"
	"github.com/Findryankp/mvcGo/apps/features/users"
	"github.com/Findryankp/mvcGo/apps/middlewares"
	"github.com/Findryankp/mvcGo/apps/utils"
)

func ExcecutionArgument(arg []string) {
	if arg[1] == "features" || arg[1] == "-f" {
		CommandFeature(arg)
	} else if arg[1] == "init" {
		CommandInit()
	}
}

func CommandFeature(arg []string) {
	if len(arg) != 3 {
		fmt.Println("wrong argument")
	} else {
		features.ModelCreate(arg[2])
		features.ControllerCreate(arg[2])
		features.RouteCreate(arg[2])
		features.MigrationCreate(arg[2])
	}
}

func CommandInit() {
	utils.FolderCreate()
	fmt.Println("setup folder success")

	configs.InitConfig()
	features.InitRouteCreate()
	middlewares.InitMiddlewares()
	users.InitUsers()
}
