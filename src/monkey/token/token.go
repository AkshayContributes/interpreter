package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
} 

var keywords = map[string]TokenType {
	"fn" : FUNCTION,
	"let" : LET,
	"int" : INT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"
	

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "fn"
	LET = "let"

)