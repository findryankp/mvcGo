package apps

import (
	"fmt"
	"os"
)

func AppendToFile() {
	file, err := os.OpenFile("./main.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Append text to file
	text := "This is some text that will be appended to the file.\n"
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
