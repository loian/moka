package lexer

import (
	"fmt"
	"moka/token"
	"unicode"
)

type Lexer struct {
	input        []rune // the input to parse
	position     int    //the current character position
	readPosition int    //the position of the next character to read
	ch           rune   //the curread read char
}

func (l *Lexer ) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) consumeWhitespace() {
	if unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() []rune {
	initialPosition := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input[initialPosition:l.position]
}

func (l *Lexer) readNumber() []rune {

}

func (l *Lexer) NextToken() token.Token {

	var t token.Token

	l.consumeWhitespace()

	switch l.ch {
	case '=':
		t = newToken(token.ASSIGN, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '[':
		t = newToken(token.LBRACKET, l.ch)
	case ']':
		t = newToken(token.RBRACKET, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t = token.Token{Type: token.EOF, Literal: ""}

	default:
		fmt.Println(l.ch)
		if unicode.IsLetter(l.ch) {
			t.Literal = string(l.readIdentifier())
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if unicode.IsDigit(l.ch) {
			t.Literal = string(l.readNumber())
			t.Type = token.TYPE_INT
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}


func newToken(tokenType token.TokenType, ch rune) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func NewLexer(code string) *Lexer {
	input :=  []rune(code)
	l := &Lexer{input: input}
	l.readChar()
	return l
}


