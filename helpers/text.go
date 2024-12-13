package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetStringFromTxt() []string {
	file, err := os.Open("model.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return []string{}
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)

	// Read line by line
	text := []string{}
	for scanner.Scan() {
		temp := scanner.Text()
		text = append(text, temp)
	}

	return text
}

func ExplodeStringFromTxt(stringFromTxt []string) map[int][]string {
	// Initialize the map
	result := make(map[int][]string)

	// Loop through the strings
	for index, text := range stringFromTxt {
		// Split each string by ":"
		parts := strings.Split(text, ":")
		result[index] = parts // Store the split parts in the map
	}

	// Return the resulting map
	return result
}
