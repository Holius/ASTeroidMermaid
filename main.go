package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/goforj/godump"
	"github.com/holius/asteroid_mermaid/filedata"
	"github.com/holius/asteroid_mermaid/mermaid"
	"github.com/holius/asteroid_mermaid/packagedata"
)

func main() {
	opts, err := parseFlags(flag.CommandLine.Args())
	if err != nil {
		fmt.Println(err)
		return
	}
	godump.Dump(opts)

	gm, err := filedata.GoMetadataFromDirectory(opts.LocalDir)
	if err != nil {
		panic(err)
	}

	gp, err := packagedata.GoPackageDataFromGoFileData(gm, false)
	if err != nil {
		panic(err)
	}

	mermaid := mermaid.GoPackagedataToMermaid(gp, opts.Module)
	om, err := os.Create("out.mermaid")
	if err != nil {
		panic(err)
	}
	_, err = om.Write([]byte(mermaid))
	if err != nil {
		panic(err)
	}
}

type Options struct {
	LocalDir string
	Module   string
}

func parseFlags(args []string) (Options, error) {
	var opts Options

	fs := flag.NewFlagSet("AsteroidMermaid", flag.ContinueOnError)

	fs.StringVar(&opts.LocalDir, "dir", "zarf/src/internal", "local directory of code to generate Mermaid document from")
	fs.StringVar(&opts.Module, "module", "github.com/zarf-dev/zarf", "Go module name belonging to dir (see go.mod)")

	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}

	var dirSet, moduleSet bool

	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "dir":
			dirSet = true
		case "module":
			moduleSet = true
		}
	})

	if dirSet != moduleSet {
		return Options{}, errors.New("-dir and -module must be provided together")
	}

	return opts, nil
}
