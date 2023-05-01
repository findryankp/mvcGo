package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func PackageIntall(link string) {
	cmd := exec.Command("go", "get", "-u", link)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(link, "Installed")
}
