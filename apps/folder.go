package apps

import (
	"fmt"
	"os"
)

func CreateFolderIfNotExist() {
	folderName := []string{"controllers", "models", "helpers", "configs", "routes"}
	for _, v := range folderName {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			// Create folder
			err = os.Mkdir(v, 0755)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
