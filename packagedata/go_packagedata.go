package packagedata

import (
	"fmt"
	"path"
	"slices"
	"strings"
	"unicode"

	"github.com/holius/asteroid_mermaid/filedata"
)

type GoPackageData struct {
	Name             string   `json:"name"` // primary key
	Files            []string `json:"files"`
	PublicFunctions  []string `json:"public_functions"`
	PrivateFunctions []string `json:"private_functions"`
	Imports          []string `json:"import"` // foreign key
}

func GoPackageDataFromGoFileData(gmSlice []filedata.GoMetadata, includeTests bool) (map[string]*GoPackageData, error) {
	pgMap := map[string]*GoPackageData{}

	for _, gm := range gmSlice {
		if strings.HasSuffix(gm.File, "_test.go") && !includeTests {
			continue
		}

		primaryKey := path.Dir(gm.File)
		pg, ok := pgMap[primaryKey]
		if !ok {
			pg = &GoPackageData{
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
		slices.Sort(gm.Imports)

		for _, function := range gm.Functions {
			if functionIsPrivate(function) {
				pg.PrivateFunctions = append(pg.PrivateFunctions, function)
			} else {
				pg.PublicFunctions = append(pg.PublicFunctions, function)
			}
		}

	}

	return pgMap, nil
}

func functionIsPrivate(funcName string) bool {
	dot := strings.Split(funcName, ".")
	if len(dot) > 2 {
		panic(fmt.Errorf("Unexpected funcName %s", funcName))
	}

	f := dot[len(dot)-1]
	if unicode.IsUpper(rune(f[0])) {
		return false
	}
	return true
}
