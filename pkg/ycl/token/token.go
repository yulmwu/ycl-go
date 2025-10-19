package token

type Type string

type Token struct {
	Type    Type
	Literal string
	Offset  int
	Line    int
	Column  int
}

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	IDENT    Type = "IDENT"
	NUMBER   Type = "NUMBER"
	STRING   Type = "STRING"
	TEMPLATE Type = "TEMPLATE"

	ASSIGN Type = "="
	COMMA  Type = ","

	LBRACE Type = "{"
	RBRACE Type = "}"
	LBRACK Type = "["
	RBRACK Type = "]"
	LPAREN Type = "("
	RPAREN Type = ")"
	DOT    Type = "."
	BANG   Type = "!"
	PLUS   Type = "+"
)

var keywords = map[string]Type{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

const (
	TRUE  Type = "TRUE"
	FALSE Type = "FALSE"
	NULL  Type = "NULL"
)

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
