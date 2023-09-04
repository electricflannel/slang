package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	CONTAINS = "CONTAINS"
	AND      = "AND"
	OR       = "OR"
	NOT      = "NOT"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)

var keywords = map[string]TokenType{
	"contains": CONTAINS,
	"and":      AND,
	"or":       OR,
	"not":      NOT,
	// "in":       IN,

	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupIdent(ident string) TokenType {
	lident := strings.ToLower(ident)

	if tok, ok := keywords[lident]; ok {
		return tok
	}
	return IDENT
}
