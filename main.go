package main

import (
	"fmt"
	"os"

	"github.com/scortier/go-parser/parser"
)

func main() {
	filePath := "tests/step2/valid.json" // Change this to test invalid.json

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	jsonParser := parser.NewParser(file)
	result, err := jsonParser.Parse()

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	if result {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
