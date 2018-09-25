package astparsestr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// AstParseFile return parsed *ast.File object or error
func AstParseFile(s string) (*ast.File, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", "package main"+s, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse: %v", err)
	}
	return f, nil
}

// AstParseExpr provides parsing of ast expression
func AstParseExpr(s string) ast.Expr {
	node, err := parser.ParseExpr(s)
	if err != nil {
		return &ast.BadExpr{}
	}
	return node
}
