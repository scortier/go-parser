// main.go
package main

import (
	"fmt"

	"github.com/scortier/go-parser/lexer"

	"os"

	"github.com/scortier/go-parser/parser"
)

func main() {
	input := "{}"

	tokens := lexer.Tokenize(input)
	_, err := parser.Parse(tokens)

	if err != nil {
		fmt.Println("Invalid JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
	os.Exit(0)
}
