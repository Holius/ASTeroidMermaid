package filedata

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/goforj/godump"
)

type GoMetadata struct {
	File      string   `json:"file"`
	Package   string   `json:"package"`
	Imports   []string `json:"imports"`
	Functions []string `json:"functions"`
}

func GoMetadataFromDirectory(directory string) ([]GoMetadata, error) {
	gmSlice := []GoMetadata{}

	filesProcessed := 0

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			if strings.HasSuffix(path, ".go") {
				gm, err := GoMetadataFromFile(path)
				if err != nil {
					return err
				}
				if gm == nil {
					return fmt.Errorf("recived null GoMetadata pointer when walking directory")
				}
				gmSlice = append(gmSlice, *gm)
				filesProcessed++
				if filesProcessed == 3 {
					//return fmt.Errorf("break")
				}

			}
		}
		return nil
	})

	return gmSlice, err
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
			//if x.Name.Name == "IsOptional" {
			//	godump.Dump(x)
			//	fmt.Println(getStructName(x))
			//  os.Exit(1)
			//}
			funcId := x.Name.Name
			if sn := getStructName(x); sn != "" {
				funcId = fmt.Sprintf("%s.%s", sn, x.Name.Name)
			}
			gm.Functions = append(gm.Functions, funcId)

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

func getStructName(funcDefinition *ast.FuncDecl) string {
	if funcDefinition.Recv == nil {
		return ""
	}
	if funcDefinition.Recv.List == nil {
		return ""
	}
	if len(funcDefinition.Recv.List) > 1 {
		godump.Dump(funcDefinition)
		panic(fmt.Errorf("Exptect 1 or less in List got %d", len(funcDefinition.Recv.List)))
	}
	switch t := funcDefinition.Recv.List[0].Type.(type) {
	case *ast.Ident:
		// func (s Server) Method()
		return t.Name

	case *ast.StarExpr:
		// func (s *Server) Method()
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	}
	return ""
}
