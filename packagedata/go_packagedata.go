package packagedata

import (
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/holius/asteroid_mermaid/filedata"
)

type GoPackageDatea struct {
	Name             string   `json:"name"` // primary key
	Files            []string `json:"files"`
	PublicFunctions  []string `json:"public_functions"`
	PrivateFunctions []string `json:"private_functions"`
	Imports          []string `json:"import"` // foreign key
}

func GoPackageDataFromGoMetadata(gmSlice []filedata.GoMetadata, includeTests bool) (map[string]*GoPackageDatea, error) {
	pgMap := map[string]*GoPackageDatea{}

	for _, gm := range gmSlice {
		if strings.HasSuffix(gm.File, "_test.go") && !includeTests {
			continue
		}

		primaryKey := path.Dir(gm.File)
		pg, ok := pgMap[primaryKey]
		if !ok {
			pg = &GoPackageDatea{
				Name:             gm.Package,
				Files:            []string{},
				PublicFunctions:  []string{},
				PrivateFunctions: []string{},
			}
			pgMap[primaryKey] = pg
		}

		pg.Files = append(pg.Files, gm.File)

		for _, imp := range gm.Imports {
			if !slices.Contains(pg.Imports, imp) {
				pg.Imports = append(pg.Imports, imp)
			}
		}

		for _, function := range gm.Functions {
			if unicode.IsUpper(rune(function[0])) {
				pg.PublicFunctions = append(pg.PublicFunctions, function)
			} else {
				pg.PrivateFunctions = append(pg.PrivateFunctions, function)
			}
		}

	}

	return pgMap, nil
}
