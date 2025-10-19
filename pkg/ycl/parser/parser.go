package parser

import (
	"fmt"

	"github.com/yulmwu/ycl-go/pkg/ycl/ast"
	"github.com/yulmwu/ycl-go/pkg/ycl/lexer"
	"github.com/yulmwu/ycl-go/pkg/ycl/token"
)

type Parser struct {
	l *lexer.Lexer

	cur  token.Token
	peek token.Token

	errors []error
}

func New(input string) *Parser {
	l := lexer.New(input)
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []error { return p.errors }

func (p *Parser) nextToken() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) expect(t token.Type) bool {
	if p.cur.Type == t {
		return true
	}
	p.errors = append(p.errors, fmt.Errorf("expected %s, got %s", t, p.cur.Type))
	return false
}

func (p *Parser) peekIs(t token.Type) bool { return p.peek.Type == t }
func (p *Parser) curIs(t token.Type) bool  { return p.cur.Type == t }

func (p *Parser) ParseFile() *ast.File {
	f := &ast.File{}
	for !p.curIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			f.Statements = append(f.Statements, stmt)
		} else {
			p.nextToken()
		}
		if p.curIs(token.COMMA) {
			p.nextToken()
		}
	}
	return f
}

func (p *Parser) parseStatement() ast.Statement {
	if p.curIs(token.IDENT) {
		name := p.cur.Literal
		if p.peekIs(token.ASSIGN) {
			p.nextToken() // move to =
			p.nextToken() // move to expr start
			expr := p.parseExpr()
			if p.peekIs(token.COMMA) {
				p.nextToken()
				p.nextToken()
			}
			return &ast.AssignStmt{Name: name, Value: expr}
		}
		if p.peekIs(token.LBRACE) {
			p.nextToken() // move to {
			p.nextToken() // first of body
			body := p.parseBlockBody()
			return &ast.BlockStmt{Name: name, Body: body}
		}
	}
	return nil
}

func (p *Parser) parseBlockBody() []ast.Statement {
	var stmts []ast.Statement
	for !p.curIs(token.RBRACE) && !p.curIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		} else {
			p.nextToken()
		}
		if p.curIs(token.COMMA) {
			p.nextToken()
		}
	}
	if p.curIs(token.RBRACE) {
		p.nextToken()
	}
	return stmts
}

func (p *Parser) parseExpr() ast.Expr {
	switch p.cur.Type {
	case token.STRING:
		lit := &ast.StringLit{Value: p.cur.Literal}
		p.nextToken()
		return lit
	case token.TEMPLATE:
		lit := &ast.TemplateLit{Raw: p.cur.Literal}
		p.nextToken()
		return lit
	case token.NUMBER:
		lit := &ast.NumberLit{Value: p.cur.Literal}
		p.nextToken()
		return lit
	case token.TRUE:
		b := &ast.BoolLit{Value: true}
		p.nextToken()
		return b
	case token.FALSE:
		b := &ast.BoolLit{Value: false}
		p.nextToken()
		return b
	case token.NULL:
		n := &ast.NullLit{}
		p.nextToken()
		return n
	case token.LBRACK:
		return p.parseArray()
	case token.IDENT:
		id := &ast.IdentRef{Name: p.cur.Literal}
		p.nextToken()
		return id
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token in expression: %s", p.cur.Type))
		p.nextToken()
		return &ast.NullLit{}
	}
}

func (p *Parser) parseArray() ast.Expr {
	// cur is [
	p.nextToken()
	arr := &ast.ArrayLit{}
	for !p.curIs(token.RBRACK) && !p.curIs(token.EOF) {
		elem := p.parseExpr()
		if elem != nil {
			arr.Elements = append(arr.Elements, elem)
		}
		if p.curIs(token.COMMA) {
			p.nextToken()
		}
	}
	if p.curIs(token.RBRACK) {
		p.nextToken()
	}
	return arr
}
