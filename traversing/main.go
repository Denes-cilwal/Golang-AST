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

/*
	Types of General Declarations
A general declaration in Go can be one of several types:

Import Declarations: Used to include packages in a Go source file. For example:
import "fmt"
import "os"


Constant Declarations: Used to define constants. For example:
const Pi = 3.14
const (
    StatusOK = 200
    StatusNotFound = 404
)


Type Declarations: Used to define new types. For example:
type Point struct {
    X, Y float64
}

Variable Declarations: Used to declare one or more variables. For example:
var a, b int
var (
    c string
    d = true
)
*/

/*
 (*ast.BinaryExpr), which are expressions involving two operands (like 3 + 2).
When it finds a binary expression, it prints whether the position of the binary expression is valid (binaryExpr.Pos().IsValid()).
It prints the right-hand side of the binary expression (binaryExpr.Y).
It also prints the line and column position of the binary expression within the file (fset.Position(binaryExpr.Pos()).Line and fset.Position(binaryExpr.Pos()).Column).
Comment on continue and IsValid():

The continue statements are used to skip over parts of the AST that don't match the current criteria (like if a declaration isn't a function declaration or a value specification).
The IsValid() method checks if the position of the binary expression is valid, which is important for ensuring that the program doesn't try to work with improperly parsed or non-existent code parts.

*/

/*
Traversing Declarations:
The program iterates over all top-level declarations (node.Decls) in the parsed file.
It checks if a declaration is a function declaration (*ast.FuncDecl). If it's not, it continues to the next declaration.

Finding the Main Function:
Within each function declaration, it checks if the function is the main function.
It then iterates over the statements (fn.Body.List) inside the main function.
Looking for Variable Declarations:

The code searches for variable declaration statements (*ast.DeclStmt).
For each declaration statement, it checks if it's a general declaration (*ast.GenDecl).
Analyzing Variable Specifications:

It iterates over the specifications (genDecl.Specs) in the general declaration, looking for variable specifications (*ast.ValueSpec).
For each variable specification, it looks at the values assigned to the variable.

Searching for Binary Expressions:
The key part is when it examines the values assigned to variables to find binary expressions (*ast.BinaryExpr), which are expressions involving two operands (like 3 + 2).
When it finds a binary expression, it prints whether the position of the binary expression is valid (binaryExpr.Pos().IsValid()).
It prints the right-hand side of the binary expression (binaryExpr.Y).
It also prints the line and column position of the binary expression within the file (fset.Position(binaryExpr.Pos()).Line and fset.Position(binaryExpr.Pos()).Column).
Comment on continue and IsValid():

The continue statements are used to skip over parts of the AST that don't match the current criteria (like if a declaration isn't a function declaration or a value specification).
The IsValid() method checks if the position of the binary expression is valid, which is important for ensuring that the program doesn't try to work with improperly parsed or non-existent code parts.

*/
