package lexer

import (
	"squee/pkg/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `(column_one == 1 AND column_two == 2)

	column_three == "foobar" OR column_four == "foo bar"

	column_five CONTAINS "a search string"

	column_six != NULL

	column_seven NOT 5
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.IDENT, "column_one"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.AND, "AND"},
		{token.IDENT, "column_two"},
		{token.EQ, "=="},
		{token.INT, "2"},
		{token.RPAREN, ")"},

		{token.IDENT, "column_three"},
		{token.EQ, "=="},
		{token.STRING, "foobar"},
		{token.OR, "OR"},
		{token.IDENT, "column_four"},
		{token.EQ, "=="},
		{token.STRING, "foo bar"},

		{token.IDENT, "column_five"},
		{token.CONTAINS, "CONTAINS"},
		{token.STRING, "a search string"},

		{token.IDENT, "column_six"},
		{token.NOT_EQ, "!="},
		{token.NULL, "NULL"},

		{token.IDENT, "column_seven"},
		{token.NOT, "NOT"},
		{token.INT, "5"},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
