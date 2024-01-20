package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := []byte(`
	package main
	import "fmt"
	func main(){
	var sum  int = 3 + 2
	fmt.Print(sum)
}`)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// start search for top level declaration
	//
	for _, decl := range node.Decls {

		// search for main function
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// search inside the main function body
		for _, stmt := range fn.Body.List {
			declStmt, ok := stmt.(*ast.DeclStmt)
			if !ok {
				continue
			}

			// continue with declStmt.Decl
			genDecl, ok := declStmt.Decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			// declarations have multiple specs
			// search for value spec

			for _, spec := range genDecl.Specs {
				valuespec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				// continue with valueSpec.Values
				for _, expr := range valuespec.Values {
					// search for binary expression
					binaryExpr, ok := expr.(*ast.BinaryExpr)
					if !ok {
						continue
					}

					fmt.Println(binaryExpr.Pos().IsValid())
					fmt.Println(binaryExpr.Y)
					// found it
					fmt.Printf("found binary expression at: %d:%d\n",
						fset.Position(binaryExpr.Pos()).Line,
						fset.Position(binaryExpr.Pos()).Column,
					)

				}
			}

		}

	}

	//ast.Fprint(os.Stdout, fset, node, nil)
}
