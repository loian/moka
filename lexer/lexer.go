package lexer

import "moka/token"

type Lexer struct {
	input        string // the input to parse
	position     int    //the current character position
	readPosition int    //the position of the next character to read
	ch           byte   //the curread read char
}

func (l *Lexer ) readChar() {
	if (l.readPosition >= len(l.input)) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}


func (l *Lexer) NextToken() token.Token {

	var t token.Token;

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
	}

	l.readChar()

	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token{
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}


