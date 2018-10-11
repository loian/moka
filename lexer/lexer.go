package lexer

import (
	"moka/token"
	"unicode"
)

type Lexer struct {
	input        []rune // the input to parse
	position     int    //the current character position
	readPosition int    //the position of the next character to read
	ch           rune   //the curread read char
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) consumeWhitespace() {
	for unicode.IsSpace(l.ch) {
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

func (l *Lexer) readNumber() ([]rune, token.TokenType) {
	initialPosition := l.position

	var tokenType token.TokenType = token.VAL_INT

	for unicode.IsNumber(l.ch) || l.ch == '.' {
		l.readChar()

		if l.ch == '.' {
			if tokenType == token.VAL_FLOAT {
				tokenType = token.ILLEGAL
			} else {
				tokenType = token.VAL_FLOAT
			}
		}
	}

	return l.input[initialPosition:l.position], tokenType
}

func (l *Lexer) peekChar() rune {
	if l.position < len(l.input) {
		return l.input[l.readPosition]
	}
	return 0
}

func (l *Lexer) twoCharOperator(tokenType token.TokenType, nextChar rune, literal string) (token.Token, bool) {
	if l.peekChar() == nextChar {
		t := token.Token{tokenType, literal}
		l.readChar() //consume the next char sine it is part of the == token
		return t, true
	}
	return token.Token{0, ""}, false
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.consumeWhitespace()

	switch l.ch {
	case '=':
		var ok bool
		if t, ok = l.twoCharOperator(token.EQUAL, '=', "=="); !ok {
			t = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '!':
		var ok bool
		if t, ok = l.twoCharOperator(token.NOT_EQUAL, '=', "!="); !ok {
			t = newToken(token.BANG, l.ch)
		}

	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
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
		if unicode.IsLetter(l.ch) {
			t.Literal = string(l.readIdentifier())
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if unicode.IsDigit(l.ch) || l.ch == '.' {
			runesLiteral, tokType := l.readNumber()
			t.Literal = string(runesLiteral)
			t.Type = tokType
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func NewLexer(code string) *Lexer {
	input := []rune(code)
	l := &Lexer{input: input}
	l.readChar()
	return l
}
