package parser

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/scortier/go-parser/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

type JSONData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
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
			return p.parseJSONObject()
		case lexer.TokenEOF:
			fmt.Println("Invalid JSON: EOF reached")
			return false, nil
		case lexer.TokenError:
			fmt.Printf("Invalid JSON: %s\n", token.Value)
			return false, nil
		}
	}
}

func (p *Parser) parseJSONObject() (bool, error) {
	var jsonData JSONData

	for {
		token := p.lexer.NextToken()

		switch token.Type {
		case lexer.TokenRightBrace:
			// Successfully parsed the JSON object
			jsonString, err := json.Marshal(jsonData)
			if err != nil {
				fmt.Println("Error marshalling JSON data:", err)
				return false, err
			}

			fmt.Printf("Valid JSON: %s\n", string(jsonString))
			return true, nil
		case lexer.TokenEOF:
			if jsonData.Key != "" && jsonData.Value == nil {
				fmt.Println("Invalid JSON: Value missing for key")
				return false, nil
			}
			fmt.Println("Invalid JSON: EOF reached")
			return false, nil
		case lexer.TokenError:
			fmt.Printf("Invalid JSON: %s\n", token.Value)
			return false, nil
		default:
			if token.Type == lexer.TokenError {
				fmt.Printf("Invalid JSON: %s\n", token.Value)
				return false, nil
			}

			if token.Type == lexer.TokenLeftBrace {
				fmt.Println("Invalid JSON: Nested objects are not supported in Step 2")
				return false, nil
			}

			if token.Type == lexer.TokenRightBrace {
				fmt.Println("Invalid JSON: Unexpected '}' encountered")
				return false, nil
			}

			if token.Type != lexer.TokenError {
				switch jsonData.Key {
				case "":
					jsonData.Key = token.Value
				default:
					if jsonData.Value == nil {
						if token.Value == ":" {
							// Start parsing the value
							p.parseJSONValue(&jsonData)
						} else {
							fmt.Println("Invalid JSON: Expected ':' after key")
							return false, nil
						}
					} else {
						fmt.Println("Invalid JSON: Unexpected token after value")
						return false, nil
					}
				}
			}
		}
	}
}

func (p *Parser) parseJSONValue(data *JSONData) {
	token := p.lexer.NextToken()

	switch token.Type {
	case lexer.TokenString, lexer.TokenNumber, lexer.TokenTrue, lexer.TokenFalse, lexer.TokenNull:
		data.Value = token.Value
	case lexer.TokenLeftBrace:
		fmt.Println("Invalid JSON: Nested objects are not supported in Step 3")
	default:
		fmt.Printf("Invalid JSON: Unexpected token '%s' in value\n", token.Value)
	}
}
