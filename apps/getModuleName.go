package apps

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the first line of the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	// Extract the module name from the first line
	fields := strings.Fields(line)
	if len(fields) != 2 || fields[0] != "module" {
		return "", errors.New("invalid go.mod file")
	}
	moduleName := fields[1]

	return moduleName, nil
}
