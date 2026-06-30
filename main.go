package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

func main() {
	gm, err := GoMetadataFromDirectory("zarf")
	if err != nil {
		panic(err)
	}

	gp, err := GoPackageDataFromGoMetadata(gm, false)
	if err != nil {
		panic(err)
	}

	// todo extract from go.mod
	moduleName := "github.com/zarf-dev/zarf"

	mermaid := GoPackagedataToMermaid(gp, moduleName)
	om, err := os.Create("out.mermaid")
	if err != nil {
		panic(err)
	}
	_, err = om.Write([]byte(mermaid))
	if err != nil {
		panic(err)
	}
}

func GoMetadataFromFile(file string) (*GoMetadata, error) {

	fset := token.NewFileSet() // positions are relative to fset
	// pass nil for src because we are parsing a file
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, fmt.Errorf("the AST is nil at file %s", file)
	}

	gm := GoMetadata{
		File:      file,
		Package:   f.Name.Name,
		Imports:   []string{},
		Functions: []string{},
	}

	for _, imp := range f.Imports {
		//fmt.Println(imp.Path.Value)
		gm.Imports = append(gm.Imports, imp.Path.Value[1:len(imp.Path.Value)-1])
	}

	for _, tld := range f.Decls {
		switch x := tld.(type) {

		case *ast.GenDecl:
			//fmt.Printf("%+v\n", x)
			//if x.Tok == token.IMPORT {
			//	for _, s := range x.Specs {
			//		switch x := s.(type) {
			//		case *ast.ImportSpec:
			//			gm.Imports = append(gm.Imports, x.Path.Value)
			//		default:
			//			//panic(fmt.Sprintf("%T\n", x))
			//		}
			//	}
			//}
		case *ast.FuncDecl:
			// Skip methods (those have a receiver) — we only want package-level functions.
			if x.Recv == nil {
				gm.Functions = append(gm.Functions, x.Name.Name)
			}

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

	return &gm, nil
}

func WriteJSON(v any, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}

func PrintJSONable(v any) {
	WriteJSON(v, os.Stdout)
}
