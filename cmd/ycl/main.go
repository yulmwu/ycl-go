package main

import (
	"fmt"
	"os"
	"strings"

	lib "github.com/yulmwu/ycl-go/pkg/ycl"
	"github.com/yulmwu/ycl-go/pkg/ycl/ast"
)

// Example CLI: parses a sample YCL (or a file if provided) and prints the AST.
func main() {
	var input string
	if len(os.Args) > 1 {
		b, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to read file:", err)
			os.Exit(1)
		}
		input = string(b)
	} else {
		input = sample()
	}

	file, errs := lib.Parse(input)
	if len(errs) > 0 {
		fmt.Fprintln(os.Stderr, "parse errors:")
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, " -", e)
		}
		os.Exit(1)
	}

	fmt.Println("=== Parsed AST ===")
	printFile(file)
}

func sample() string {
	return `
my_value = 1234,
my_object {
  field1 = 'YCL',
  field2 = true,
  field3 = null,
  field4 = [ 'item1', 42, 'item3' ]
},
another_value = ` + "`" + `Hello $(my_object.field1) World!` + "`" + `
`
}

func printFile(f *ast.File) {
	for _, s := range f.Statements {
		printStmt(s, 0)
	}
}

func printStmt(s ast.Statement, indent int) {
	pad := strings.Repeat("  ", indent)
	switch n := s.(type) {
	case *ast.AssignStmt:
		fmt.Printf("%sAssign %s = ", pad, n.Name)
		printExpr(n.Value, 0)
		fmt.Println()
	case *ast.BlockStmt:
		fmt.Printf("%sBlock %s {\n", pad, n.Name)
		for _, inner := range n.Body {
			printStmt(inner, indent+1)
		}
		fmt.Printf("%s}\n", pad)
	default:
		fmt.Printf("%s<unknown stmt %T>\n", pad, s)
	}
}

func printExpr(e ast.Expr, indent int) {
	switch v := e.(type) {
	case *ast.StringLit:
		fmt.Printf("'"+"%s"+"'", v.Value)
	case *ast.TemplateLit:
		fmt.Printf("`"+"%s"+"`", v.Raw)
	case *ast.NumberLit:
		fmt.Printf("%s", v.Value)
	case *ast.BoolLit:
		if v.Value {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
	case *ast.NullLit:
		fmt.Print("null")
	case *ast.ArrayLit:
		fmt.Print("[")
		for i, el := range v.Elements {
			if i > 0 {
				fmt.Print(", ")
			}
			printExpr(el, indent+1)
		}
		fmt.Print("]")
	case *ast.IdentRef:
		fmt.Print(v.Name)
	default:
		fmt.Printf("<unknown expr %T>", e)
	}
}
