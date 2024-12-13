package helpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func CamelToSnake(input string) string {
	// Find positions where a lowercase letter is followed by an uppercase letter
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	// Replace matches with a lowercase letter followed by an underscore and the uppercase letter
	result := re.ReplaceAllString(input, "${1}_${2}")
	// Convert the entire string to lowercase
	return strings.ToLower(result)
}
