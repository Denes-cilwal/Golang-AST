package theory

// token is a set of lexical tokens of the go programming language
// the token package has two other things

type Position struct {
	filename string
	Offset   int
	Line     int
	Column   int
}

type file struct {
}

type fileSet struct {
	// multiple files
}

//  how do we get tokens ?
/*
	package scanner implements a scanner for go source text. It takes a []byte as source which can then be tokenized through a repeated calls to scan method

*/

// how do we scan ?
// - go/scanner gives you position, literal, tokens

// how to get ast
// - the parser use go/scanner to get the tokens and create the type for the ast

// how do we parse ?

// A value Spec node represent a constant or variable declaration[const or varspec]
// func (s *ast.ValueSpec) End() token.Pos

var sum int = 3 + 2

/*
				*ast.ValueSpec [if we put all them together? we get value spec]
						|
	sum         int
	|            |
*ast.Ident *ast.Ident        ast.BinaryExpression (+)
								|          |
								3           2
							ast.BasicLit  ast.BasicLit



*/

// how to traverse the tree ?
