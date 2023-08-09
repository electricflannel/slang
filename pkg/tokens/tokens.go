package tokens

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	AND   = "AND"
	OR    = "OR"
	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)

var keywords = map[string]TokenType{
	"and":   AND,
	"or":    OR,
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
