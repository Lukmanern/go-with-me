package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// Define the filename to analyze
	filename := "example.go"

	// Initialize the file set
	fset := token.NewFileSet()

	// Parse the file
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize the visitor
	visitor := &unusedVariableVisitor{}

	// Traverse the AST
	ast.Walk(visitor, node)

	// Print any unused variables
	fmt.Println("Unused variables:")
	for _, v := range visitor.unusedVariables {
		fmt.Printf("%s:%d:%d %s\n", filename, fset.Position(v.Pos()).Line, fset.Position(v.Pos()).Column, v.Name)
	}
}

type unusedVariableVisitor struct {
	unusedVariables []*ast.Ident
}

func (v *unusedVariableVisitor) Visit(node ast.Node) ast.Visitor {
	// Check for variable declarations
	if decl, ok := node.(*ast.ValueSpec); ok {
		// Check if the variable is unused
		if len(decl.Names) == 1 && decl.Names[0].IsExported() && len(decl.Names[0].Name) > 1 {
			for _, use := range decl.Names[0].Uses {
				if use != decl.Names[0] {
					return v
				}
			}
			v.unusedVariables = append(v.unusedVariables, decl.Names[0])
		}
	}
	return v
}
