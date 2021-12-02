package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // illegal token
	EOF     = "EOF"     // end of file

	// literals
	IDENT = "IDENT" // identifier, ie foo
	INT   = "INT"   // integer, eg: 45

	// operators
	ASSIGN   = "=" // self explanatory
	PLUS     = "+" // not in the sense of numeric addition
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";" // C is <3
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{" // C is life
	RBRACE    = "}"

	//keywords
	FUNCTION         = "FUNCTION"
	LET              = "LET"
	TRUE             = "TRUE"
	FALSE            = "FALSE"
	IF               = "IF"
	ELSE             = "ELSE"
	RETURN           = "RETURN"
	FLOATING_INTEGER = "FLOATING_INTEGER"
)

var keywords = map[string]TokenType{
	"fn":               FUNCTION,
	"let":              LET,
	"true":             TRUE,
	"false":            FALSE,
	"if":               IF,
	"else":             ELSE,
	"return":           RETURN,
	"floating_integer": FLOATING_INTEGER,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

// signing off
