package ast

type Node interface {
	node()
}

type File struct {
	Statements []Statement
}

func (*File) node() {}

type Statement interface {
	Node
	statement()
}

type AssignStmt struct {
	Name  string
	Value Expr
}

func (*AssignStmt) node()      {}
func (*AssignStmt) statement() {}

type BlockStmt struct {
	Name string
	Body []Statement
}

func (*BlockStmt) node()      {}
func (*BlockStmt) statement() {}

type Expr interface {
	Node
	expr()
}

type StringLit struct{ Value string }

func (*StringLit) node() {}
func (*StringLit) expr() {}

type TemplateLit struct{ Raw string }

func (*TemplateLit) node() {}
func (*TemplateLit) expr() {}

type NumberLit struct{ Value string }

func (*NumberLit) node() {}
func (*NumberLit) expr() {}

type BoolLit struct{ Value bool }

func (*BoolLit) node() {}
func (*BoolLit) expr() {}

type NullLit struct{}

func (*NullLit) node() {}
func (*NullLit) expr() {}

type ArrayLit struct{ Elements []Expr }

func (*ArrayLit) node() {}
func (*ArrayLit) expr() {}

type IdentRef struct{ Name string }

func (*IdentRef) node() {}
func (*IdentRef) expr() {}
