package main

// src is the input for which we want to print the AST.
//	src := `
//package main
//import "fmt"
//func main() {
//	fmt.Println("Hello, World!")
//}
//`
//
//	// Create the AST by parsing src.

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to inspect the AST.
	src := `
package p
import "fmt"
func Stuff() {
	fmt.Println("whatever")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	for _, tld := range f.Decls {
		fmt.Printf("%+v\n", tld)
		switch x := tld.(type) {

		default:
			fmt.Printf("%T\n", x)
			//case *ast.BasicLit:
			//	fmt.Println("BasicLit")
			//	s = x.Value
			//case *ast.Ident:
			//	fmt.Println("Ident")
			//	s = x.Name
			//case *ast.File:
			//	fmt.Println("File")
			//	s = x.Name.Name
		}
	}

	fset = token.NewFileSet() // positions are relative to fset
	f, err = parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}
