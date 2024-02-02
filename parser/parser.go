package parser

import (
	"fmt"
	"io"

	"github.com/scortier/go-parser/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		lexer: lexer.NewLexer(reader),
	}
}

func (p *Parser) Parse() (bool, error) {
	for {
		token := p.lexer.NextToken()

		switch token.Type {
		case lexer.TokenLeftBrace:
			fmt.Println("Valid JSON")
			return true, nil
		case lexer.TokenEOF:
			fmt.Println("Invalid JSON: EOF reached")
			return false, nil
		case lexer.TokenError:
			fmt.Printf("Invalid JSON: %s\n", token.Value)
			return false, nil
		}
	}
}
