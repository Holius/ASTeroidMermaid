package mermaid

import (
	"fmt"
	"path"
	"slices"
	"strings"

	"github.com/holius/asteroid_mermaid/packagedata"
)

func GoPackagedataToMermaid(pdMap map[string]*packagedata.GoPackageData, moduleName string) string {
	moduleName = path.Dir(moduleName)
	mermaid := "graph TD\n"

	// getting key order makes output consistent
	keys := []string{}
	for k := range pdMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for i := range len(keys) {
		primaryKey := keys[i]
		fmt.Println(primaryKey)
		pd := pdMap[primaryKey]
		mermaid += fmt.Sprintf("  subgraph %s\n", primaryKey)
		for _, pf := range pd.PublicFunctions {
			mermaid += fmt.Sprintf("  %s\n", pf)
		}
		mermaid += "  end\n"
		for _, foreignKey := range pd.Imports {
			fk, found := strings.CutPrefix(foreignKey, moduleName)
			if found {
				mermaid += fmt.Sprintf("  %s --> %s\n", primaryKey, strings.TrimLeft(fk, "/"))
			}
		}
	}
	return mermaid
}
