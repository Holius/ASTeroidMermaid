package main

import (
	"encoding/json"
	"io"
	"os"
	"reflect"

	"github.com/holius/asteroid_mermaid/filedata"
	"github.com/holius/asteroid_mermaid/mermaid"
	"github.com/holius/asteroid_mermaid/packagedata"
)

func main() {
	gm, err := filedata.GoMetadataFromDirectory("zarf")
	if err != nil {
		panic(err)
	}

	gp, err := packagedata.GoPackageDataFromGoMetadata(gm, false)
	if err != nil {
		panic(err)
	}

	// todo extract from go.mod
	moduleName := "github.com/zarf-dev/zarf"

	mermaid := mermaid.GoPackagedataToMermaid(gp, moduleName)
	om, err := os.Create("out.mermaid")
	if err != nil {
		panic(err)
	}
	_, err = om.Write([]byte(mermaid))
	if err != nil {
		panic(err)
	}
}

func getName(v any) string {
	if v == nil {
		return ""
	}
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}

func WriteJSON(v any, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(v)
}

func PrintJSONable(v any) {
	WriteJSON(v, os.Stdout)
}
