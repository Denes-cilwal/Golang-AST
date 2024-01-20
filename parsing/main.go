package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	src := []byte(`package main; var sum int =  2 + 3;`)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Fprint(os.Stdout, fset, node, nil)
}

/*
a File Set in Go can track positions within a single source code file as well.
While File Sets are often used to manage positions across multiple files in a program, they are equally capable of tracking positions within a single file.
When you create a File Set using token.NewFileSet(), it starts as an empty container.
As you parse a source code file using parser.ParseFile and provide the File Set as an argument,
the parser associates positions within that file with the File Set.
This allows you to track positions accurately within the single source code file being processed.

we convert the source code to a byte slice and set it for the file in the File Set using file.SetBytes([]byte(sourceCode)).
This step is necessary because the parser.ParseFile function expects the content of the file to be provided as a byte slice.

Here's why we convert the source code to bytes:
File Content Format: In Go, source code files are typically read from the file system as binary data,
represented as a sequence of bytes. This is the common way to work with file content in most programming languages.
*/

/*



 */
