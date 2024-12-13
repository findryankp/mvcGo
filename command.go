package mvcGo

import (
	"errors"
	"fmt"
	"os"

	"github.com/Findryankp/mvcGo/apps"
)

func Init() error {
	var argsRaw = os.Args
	if len(argsRaw) >= 2 {
		Command(argsRaw)
		return errors.New("run command finish")
	}

	return nil
}

func Command(argsRaw []string) {
	args := []string{
		"features", "init", "featuresTxt", "-i", "-f", "-ft",
	}

	flag := false
	for _, v := range args {
		if v == argsRaw[1] {
			apps.ExcecutionArgument(argsRaw)
			flag = true
		}
	}

	if !flag {
		fmt.Println(`1. "init" for init project`)
		fmt.Println(`2. "features or -f" for create feature`)
		fmt.Println(`3. "featuresTxt or -ft" for create feature from Model.TXT`)
	}
}
