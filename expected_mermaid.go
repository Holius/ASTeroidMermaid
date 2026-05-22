package main

// the belodw data was modified from actual source code of Zarf
var mermaidInput = map[string]*GoPackagaedata{
	"zarf": {
		Files:           []string{"zarf/main.go"},
		Name:            "main",
		PublicFunctions: []string{},
		PrivateFunctions: []string{
			"main",
		},
		Imports: []string{
			"context",
			"os",
			"os/signal",
			"syscall",
			"github.com/zarf-dev/zarf/src/cmd",
			"github.com/zarf-dev/zarf/src/config",
		},
	},
	"zarf/src/config": {
		Files: []string{
			"zarf/src/config/template.go",
		},
		Name: "template",
		PublicFunctions: []string{
			"NewObjects",
			"WithValues",
			"WithMetadata",
			"WithBuild",
			"WithConstants",
			"WithVariables",
			"WithPackage",
			"Apply",
			"ApplyToFile",
		},
		PrivateFunctions: []string{},
		Imports: []string{
			"bytes",
			"context",
			"encoding/json",
			"fmt",
			"maps",
			"os",
			"path/filepath",
			"strings",
			"text/template",
			"time",
			"github.com/BurntSushi/toml",
			"github.com/Masterminds/sprig/v3",
			"github.com/goccy/go-yaml",
			"github.com/zarf-dev/zarf/src/api/v1alpha1",
		},
	},
	"zarf/src/cmd": {
		Files: []string{
			"zarf/src/cmd/dns.go",
		},
		Name: "dns",
		PublicFunctions: []string{
			"IsServiceURL",
			"ParseServiceURL",
			"IsLocalhost",
		},
		PrivateFunctions: []string{},
		Imports: []string{
			"errors",
			"fmt",
			"net",
			"net/url",
			"regexp",
			"strconv",
			"strings",
		},
	},
	"zarf/src/internal/api/v1alpha1": {
		Files: []string{
			"zarf/src/internal/api/v1alpha1/validate.go",
		},
		Name: "v1alpha1",
		PublicFunctions: []string{
			"ValidatePackage",
		},
		PrivateFunctions: []string{
			"validateActions",
			"hasSetVariables",
			"hasTemplating",
			"validateActionSet",
			"validateAction",
			"validateReleaseName",
			"validateChart",
			"validateManifest",
		},
		Imports: []string{
			"errors",
			"fmt",
			"strings",
			"github.com/zarf-dev/zarf/src/api/v1alpha1",
			"k8s.io/apimachinery/pkg/util/validation",
		},
	},
	"zarf/src/internal/pkgcfg": {
		Files: []string{
			"zarf/src/internal/pkgcfg/pkgcfg.go",
		},
		Name: "pkgcfg",
		PublicFunctions: []string{
			"Parse",
			"ParseMultiDoc",
		},
		PrivateFunctions: []string{
			"decodeV1Alpha1",
			"applyV1Alpha1Migrations",
			"handlerFor",
			"apiVersionFromNode",
			"parseZarfYAMLDocs",
			"filterEmptyDocs",
			"migrateDeprecated",
			"migrateScriptsToActions",
			"migrateSetVariableToSetVariables",
			"clearSetVariables",
		},
		Imports: []string{
			"context",
			"errors",
			"fmt",
			"math",
			"slices",
			"github.com/goccy/go-yaml",
			"github.com/goccy/go-yaml/ast",
			"github.com/goccy/go-yaml/parser",
			"github.com/zarf-dev/zarf/src/api/v1alpha1",
		},
	},
	"zarf/src/internal/git": {
		Files: []string{
			"zarf/src/internal/git/repository.go",
		},
		Name: "git",
		PublicFunctions: []string{
			"Open",
			"Clone",
			"Path",
			"Push",
		},
		PrivateFunctions: []string{
			"checkoutRefAsBranch",
		},
		Imports: []string{
			"context",
			"errors",
			"fmt",
			"os",
			"path/filepath",
			"strings",
			"github.com/go-git/go-git/v5",
			"github.com/go-git/go-git/v5/config",
			"github.com/go-git/go-git/v5/plumbing",
			"github.com/go-git/go-git/v5/plumbing/object",
			"github.com/go-git/go-git/v5/plumbing/transport",
			"github.com/go-git/go-git/v5/plumbing/transport/http",
			"github.com/zarf-dev/zarf/src/cmd",
			"github.com/zarf-dev/zarf/src/config",
			"github.com/zarf-dev/zarf/src/api/v1alpha1",
		},
	},
	"zarf-dev/zarf/src/api/v1alpha1": {
		Files: []string{
			"zarf/src/internal/git/repository.go",
		},
		Name: "git",
		PublicFunctions: []string{
			"Source",
		},
		PrivateFunctions: []string{
			"checkoutRefAsBranch",
		},
		Imports: []string{
			"context",
		},
	},
}

var mermaidExpected = `graph TD
  subgraph zarf
  end
  zarf --> zarf/src/cmd
  zarf --> zarf/src/config
  subgraph zarf-dev/zarf/src/api/v1alpha1
  Source
  end
  subgraph zarf/src/cmd
  IsServiceURL
  ParseServiceURL
  IsLocalhost
  end
  subgraph zarf/src/config
  NewObjects
  WithValues
  WithMetadata
  WithBuild
  WithConstants
  WithVariables
  WithPackage
  Apply
  ApplyToFile
  end
  zarf/src/config --> zarf/src/api/v1alpha1
  subgraph zarf/src/internal/api/v1alpha1
  ValidatePackage
  end
  zarf/src/internal/api/v1alpha1 --> zarf/src/api/v1alpha1
  subgraph zarf/src/internal/git
  Open
  Clone
  Path
  Push
  end
  zarf/src/internal/git --> zarf/src/cmd
  zarf/src/internal/git --> zarf/src/config
  zarf/src/internal/git --> zarf/src/api/v1alpha1
  subgraph zarf/src/internal/pkgcfg
  Parse
  ParseMultiDoc
  end
  zarf/src/internal/pkgcfg --> zarf/src/api/v1alpha1
`
