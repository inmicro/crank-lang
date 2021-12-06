package lexer

import "crank/token"

type Lexer struct {
	input        string
	position     int  // position in input
	readPosition int  // reading position in input
	ch           byte // current char under examination
}

// takes in input as string and gives lexer atrributes
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// reads character, position and increments position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // either EOF or no input
	} else {
		l.ch = l.input[l.readPosition] // read character to byte
	}

	l.position = l.readPosition // set position to readposition
	l.readPosition += 1         //increment readposition to move onto next character. Useful for peek

}

// reads and identifies tokens.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// reads identifiers and keywords
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// for numbers
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// for ILLEGAL tokens
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// generates a new token. Useful in v1 repl. Helper function for NextToken
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// reads variable names
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// reads and returns letters+symbols. Works for small, capital letters and _, !.
// Note: reads small letters and capital letters as different, and hence, caps and smalls
// will go as diff idents
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && 'Z' <= ch || ch == '_' || ch == '!'

}

// self explanotory Skips spaces, tabs, carriage returns and newlines.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

// reads a number, using the the isDigit function.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Checks for digits, and returns between 0 & 9. Helps in numbers
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'

}

// Looks ahead for the next byte, ie at readposition, returns it. Especially useful for == and !=
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// TODO: IMPLEMENT UNICODE
// TODO: IMPLEMENT FLOATS
// TODO: IMPLEMENT HEX NOTATION
