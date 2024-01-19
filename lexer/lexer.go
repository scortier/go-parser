// lexer/lexer.go
package lexer

type TokenType int

const (
	ObjectStart TokenType = iota
	ObjectEnd
	EOF
	Invalid
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) []Token {
	tokens := make([]Token, 0)

	for _, char := range input {
		switch char {
		case '{':
			tokens = append(tokens, Token{ObjectStart, "{"})
		case '}':
			tokens = append(tokens, Token{ObjectEnd, "}"})
		}
	}

	tokens = append(tokens, Token{EOF, ""})
	return tokens
}
