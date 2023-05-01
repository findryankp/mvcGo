package apps

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLineNumber(fileName string, text string) int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, text) {
			break
		}
		lineNumber++
	}
	return lineNumber
}
