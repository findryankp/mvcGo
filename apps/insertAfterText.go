package apps

import (
	"bufio"
	"fmt"
	"os"
)

func InsertAfterText(fileName, content string, lineNumber int) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Insert text in line 3
	if len(lines) >= 3 {
		lines[lineNumber] = content
	}

	// Write the modified lines back to the file
	file.Truncate(0) // Clear the file
	file.Seek(0, 0)  // Reset the file offset
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}
