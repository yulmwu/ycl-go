package ycl

import (
	"io"
	"os"

	"github.com/yulmwu/ycl-go/pkg/ycl/ast"
	"github.com/yulmwu/ycl-go/pkg/ycl/parser"
)

func Hello() string {
	return "test"
}

func Parse(input string) (*ast.File, []error) {
	p := parser.New(input)
	file := p.ParseFile()
	return file, p.Errors()
}

func ParseFile(path string) (*ast.File, []error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, []error{err}
	}
	return Parse(string(b))
}

func ParseReader(r io.Reader) (*ast.File, []error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, []error{err}
	}
	return Parse(string(b))
}
