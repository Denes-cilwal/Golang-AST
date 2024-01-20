package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	src := `
package main

import "fmt"

func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println(Add(3, 2))
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		// Check if the node is a function declaration
		if fn, ok := n.(*ast.FuncDecl); ok {
			log.Printf("Found a function: %s", fn.Name.Name)
			log.Printf("found a function: %T", fn.Type)
		}
		return true // return true to continue the traversal
	})

	printer.Fprint(os.Stdout, fset, file) // go-printer nodes print ast nodes in the human read able form
}

/*
Parse the Source Code: First, you need to parse your Go source code to create an AST. You can use the parser.ParseFile function for this.

Define a Visitor Function: This function should have the signature func(ast.Node) bool.
It's called for each node in the AST. The function can inspect or modify the node.
If it returns true, ast.Inspect continues to traverse the tree into the children of this node.
If it returns false, ast.Inspect skips the children of this node.

Call ast.Inspect: Apply ast.Inspect to the root of the AST and pass your visitor function.
It will traverse the AST and apply your function to each node.

*/
