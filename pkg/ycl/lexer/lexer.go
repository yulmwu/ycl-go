package lexer

import (
	"unicode"
	"unicode/utf8"

	"github.com/yulmwu/ycl-go/pkg/ycl/token"
)

type Lexer struct {
	input string
	pos   int
	read  int
	ch    rune
	width int
	line  int
	col   int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, col: 0}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespaceAndComments()

	tok := token.Token{Offset: l.pos, Line: l.line, Column: l.col}

	switch l.ch {
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	case '=':
		tok = l.emit(token.ASSIGN, "=")
		l.readRune()
	case ',':
		tok = l.emit(token.COMMA, ",")
		l.readRune()
	case '{':
		tok = l.emit(token.LBRACE, "{")
		l.readRune()
	case '}':
		tok = l.emit(token.RBRACE, "}")
		l.readRune()
	case '[':
		tok = l.emit(token.LBRACK, "[")
		l.readRune()
	case ']':
		tok = l.emit(token.RBRACK, "]")
		l.readRune()
	case '(':
		tok = l.emit(token.LPAREN, "(")
		l.readRune()
	case ')':
		tok = l.emit(token.RPAREN, ")")
		l.readRune()
	case '.':
		tok = l.emit(token.DOT, ".")
		l.readRune()
	case '!':
		tok = l.emit(token.BANG, "!")
		l.readRune()
	case '+':
		tok = l.emit(token.PLUS, "+")
		l.readRune()
	case '\'':
		lit := l.readString('\'')
		tok = l.emit(token.STRING, lit)
	case '`':
		lit := l.readString('`')
		tok = l.emit(token.TEMPLATE, lit)
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			t := token.LookupIdent(ident)
			tok = l.emit(t, ident)
		} else if isDigit(l.ch) {
			num := l.readNumber()
			tok = l.emit(token.NUMBER, num)
		} else {
			lit := string(l.ch)
			tok = l.emit(token.ILLEGAL, lit)
			l.readRune()
		}
	}
	return tok
}

func (l *Lexer) emit(t token.Type, lit string) token.Token {
	return token.Token{Type: t, Literal: lit, Offset: l.pos, Line: l.line, Column: l.col}
}

func (l *Lexer) readRune() {
	if l.read >= len(l.input) {
		l.pos = l.read
		l.ch = 0
		l.width = 0
		return
	}
	r, w := utf8.DecodeRuneInString(l.input[l.read:])
	l.pos = l.read
	l.read += w
	l.ch = r
	l.width = w
	if r == '\n' {
		l.line++
		l.col = 0
	} else {
		l.col++
	}
}

func (l *Lexer) peekRune() rune {
	if l.read >= len(l.input) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.read:])
	return r
}

func (l *Lexer) skipWhitespaceAndComments() {
	for {
		if isWhitespace(l.ch) {
			l.readRune()
			continue
		}
		if l.ch == '/' && l.peekRune() == '/' {
			for l.ch != 0 && l.ch != '\n' {
				l.readRune()
			}
			continue
		}
		break
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.pos
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readRune()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readNumber() string {
	start := l.pos
	for isDigit(l.ch) {
		l.readRune()
	}
	return l.input[start:l.pos]
}

func (l *Lexer) readString(quote rune) string {
	l.readRune()
	start := l.pos
	for l.ch != 0 && l.ch != quote {
		if l.ch == '\\' {
			l.readRune()
			if l.ch == 0 {
				break
			}
		}
		l.readRune()
	}
	lit := l.input[start:l.pos]
	if l.ch == quote {
		l.readRune()
	}
	return lit
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func isLetter(r rune) bool {
	return r == '_' || r == '$' || unicode.IsLetter(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
