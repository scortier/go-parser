// parser/parser.go
package parser

import (
	"fmt"

	"github.com/scortier/go-parser/lexer"
)

type Node struct {
	Type     string
	Value    string
	Children []*Node
}

func Parse(tokens []lexer.Token) (*Node, error) {
	root := &Node{Type: "Root", Value: ""}

	for tokens[0].Type != lexer.EOF {
		token := tokens[0]

		switch token.Type {
		case lexer.ObjectStart:
			objectNode := &Node{Type: "Object", Value: token.Value}
			root.Children = append(root.Children, objectNode)
			tokens = tokens[1:]
		case lexer.ObjectEnd:
			if len(root.Children) == 0 || root.Children[len(root.Children)-1].Type != "Object" {
				return nil, fmt.Errorf("Syntax error: Unexpected '}'")
			}
			tokens = tokens[1:]
		default:
			return nil, fmt.Errorf("Syntax error: Unexpected token %s", token.Value)
		}
	}

	return root, nil
}
