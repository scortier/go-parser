package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scortier/go-parser/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the path to the JSON file: ")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove newline character from the input
	filePath = filepath.Clean(filePath[:len(filePath)-1])

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", filePath, err)
		return
	}
	defer file.Close()

	fmt.Printf("\nRunning parser for file: %s\n", filePath)

	jsonParser := parser.NewParser(file)
	result, err := jsonParser.Parse()

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if result {
		fmt.Println("Parser reported JSON as valid.")
	} else {
		fmt.Println("Parser reported JSON as invalid.")
	}
}
