package lexer

import (
	"bufio"
	"io"
)

type TokenType int

const (
	TokenError TokenType = iota
	TokenLeftBrace
	TokenRightBrace
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) NextToken() Token {
	char, _, err := l.reader.ReadRune()

	if err != nil {
		if err == io.EOF {
			return Token{Type: TokenEOF, Value: ""}
		}
		return Token{Type: TokenError, Value: err.Error()}
	}

	switch char {
	case '{':
		return Token{Type: TokenLeftBrace, Value: string(char)}
	case '}':
		return Token{Type: TokenRightBrace, Value: string(char)}
	default:
		return Token{Type: TokenError, Value: "Invalid token"}
	}
}
