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
	// Check for variable usages
	if ident, ok := node.(*ast.Ident); ok {
		// Check if the variable is unused
		if len(ident.Name) > 1 && ident.IsExported() {
			// Check if the variable has an associated object
			if ident.Obj != nil && ident.Obj.Kind == ast.Var {
				// Iterate over the usages of
				// the variable's definition
				for _, use := range ident.Obj.Decl.(*ast.ValueSpec).Names {
					if use != ident {
						// The variable is used elsewhere,
						// so it's not unused
						return v
					}
				}
				// The variable is unused
				v.unusedVariables = append(v.unusedVariables, ident)
			}
		}
	}
	return v
}
