# ASTeroid Mermaid

This app maps public/exported functions in a package
to other intra packages that it imports
- "intra package" means a package that exists within the same module

Below is expected test output from fake package data inspired by Zarf source code:
```mermaid
graph TD
  subgraph zarf/src/internal/agent
  StartWebhook
  StartHTTPProxy
  end
  zarf/src/internal/agent --> zarf/src/internal/agent/hooks
  zarf/src/internal/agent --> zarf/src/internal/agent/http
  zarf/src/internal/agent --> zarf/src/internal/agent/http/admission
  zarf/src/internal/agent --> zarf/src/pkg/cluster
  zarf/src/internal/agent --> zarf/src/pkg/logger
  subgraph zarf/src/internal/agent/hooks
  NewApplicationMutationHook
  NewApplicationSetMutationHook
  NewAppProjectMutationHook
  NewRepositorySecretMutationHook
  NewGitRepositoryMutationHook
  NewHelmRepositoryMutationHook
  NewOCIRepositoryMutationHook
  NewPodMutationHook
  end
  zarf/src/internal/agent/hooks --> zarf/src/config/lang
  zarf/src/internal/agent/hooks --> zarf/src/internal/agent/operations
  zarf/src/internal/agent/hooks --> zarf/src/pkg/cluster
  zarf/src/internal/agent/hooks --> zarf/src/pkg/logger
  zarf/src/internal/agent/hooks --> zarf/src/pkg/state
  zarf/src/internal/agent/hooks --> zarf/src/pkg/transform
  zarf/src/internal/agent/hooks --> zarf/src/pkg/pki
  zarf/src/internal/agent/hooks --> zarf/src/pkg/images
  zarf/src/internal/agent/hooks --> zarf/src/config
  subgraph zarf/src/internal/agent/http
  ProxyHandler
  end
  zarf/src/internal/agent/http --> zarf/src/pkg/cluster
  zarf/src/internal/agent/http --> zarf/src/pkg/logger
  zarf/src/internal/agent/http --> zarf/src/pkg/state
  zarf/src/internal/agent/http --> zarf/src/pkg/transform
  subgraph zarf/src/internal/agent/http/admission
  NewHandler
  Handler.Serve
  end
  zarf/src/internal/agent/http/admission --> zarf/src/config/lang
  zarf/src/internal/agent/http/admission --> zarf/src/internal/agent/operations
  zarf/src/internal/agent/http/admission --> zarf/src/pkg/logger
  subgraph zarf/src/internal/agent/operations
  Hook.Execute
  AddPatchOperation
  RemovePatchOperation
  ReplacePatchOperation
  CopyPatchOperation
  MovePatchOperation
  end
  zarf/src/internal/agent/operations --> zarf/src/config/lang
  subgraph zarf/src/internal/api/v1alpha1
  ValidatePackage
  end
  zarf/src/internal/api/v1alpha1 --> zarf/src/api/v1alpha1
  subgraph zarf/src/internal/dns
  IsServiceURL
  ParseServiceURL
  IsLocalhost
  end
  subgraph zarf/src/internal/git
  ParseRef
  Open
  Clone
  Repository.Path
  Repository.Push
  end
  zarf/src/internal/git --> zarf/src/pkg/utils/exec
  zarf/src/internal/git --> zarf/src/pkg/logger
  zarf/src/internal/git --> zarf/src/pkg/transform
  zarf/src/internal/git --> zarf/src/pkg/utils
  subgraph zarf/src/internal/gitea
  NewClient
  Client.DoRequest
  Client.CreateReadOnlyUser
  Client.UpdateGitUser
  Client.CreatePackageRegistryToken
  Client.AddReadOnlyUserToRepository
  end
  subgraph zarf/src/internal/healthchecks
  Run
  WaitForReadyRuntime
  WaitForReady
  NewImmediateWatcher
  ImmediateWatcher.Watch
  end
  zarf/src/internal/healthchecks --> zarf/src/api/v1alpha1
  subgraph zarf/src/internal/packager/helm
  InstallOrUpgradeChart
  RemoveChart
  UpdateReleaseValues
  LoadChartData
  ChartFromZarfManifest
  StandardName
  StandardValuesName
  Destroy
  FindAnnotatedImagesForChart
  renderer.Run
  PackageChart
  PackageChartFromLocalFiles
  PackageChartFromGit
  DownloadPublishedChart
  DownloadChartFromGitToTemp
  TemplateChart
  templateRenderer.Run
  UpdateZarfRegistryValues
  UpdateZarfAgentValues
  end
  zarf/src/internal/packager/helm --> zarf/src/api/v1alpha1
  zarf/src/internal/packager/helm --> zarf/src/pkg/cluster
  zarf/src/internal/packager/helm --> zarf/src/pkg/logger
  zarf/src/internal/packager/helm --> zarf/src/pkg/state
  zarf/src/internal/packager/helm --> zarf/src/pkg/variables
  zarf/src/internal/packager/helm --> zarf/src/internal/healthchecks
  zarf/src/internal/packager/helm --> zarf/src/internal/packager/template
  zarf/src/internal/packager/helm --> zarf/src/config
  zarf/src/internal/packager/helm --> zarf/src/types
  zarf/src/internal/packager/helm --> zarf/src/config/lang
  zarf/src/internal/packager/helm --> zarf/src/internal/git
  zarf/src/internal/packager/helm --> zarf/src/pkg/transform
  zarf/src/internal/packager/helm --> zarf/src/pkg/utils
  subgraph zarf/src/internal/packager/kustomize
  Build
  end
  subgraph zarf/src/internal/packager/requirements
  VersionRequirementsError.Error
  ValidateVersionRequirements
  end
  zarf/src/internal/packager/requirements --> zarf/src/api/v1alpha1
  zarf/src/internal/packager/requirements --> zarf/src/config
  subgraph zarf/src/internal/packager/template
  GetZarfVariableConfig
  GetZarfTemplates
  end
  zarf/src/internal/packager/template --> zarf/src/pkg/state
  zarf/src/internal/packager/template --> zarf/src/api/v1alpha1
  zarf/src/internal/packager/template --> zarf/src/config
  zarf/src/internal/packager/template --> zarf/src/pkg/interactive
  zarf/src/internal/packager/template --> zarf/src/pkg/logger
  zarf/src/internal/packager/template --> zarf/src/pkg/utils
  zarf/src/internal/packager/template --> zarf/src/pkg/variables
  subgraph zarf/src/internal/pkgcfg
  Parse
  ParseMultiDoc
  end
  zarf/src/internal/pkgcfg --> zarf/src/api/v1alpha1
  zarf/src/internal/pkgcfg --> zarf/src/pkg/logger
  subgraph zarf/src/internal/split
  SplitFile
  ReassembleFile
  end
  zarf/src/internal/split --> zarf/src/pkg/logger
  subgraph zarf/src/internal/template
  NewObjects
  Objects.WithValues
  Objects.WithMetadata
  Objects.WithBuild
  Objects.WithConstants
  Objects.WithVariables
  Objects.WithPackage
  Apply
  ApplyToFile
  end
  zarf/src/internal/template --> zarf/src/api/v1alpha1
  zarf/src/internal/template --> zarf/src/pkg/logger
  zarf/src/internal/template --> zarf/src/pkg/value
  zarf/src/internal/template --> zarf/src/pkg/variables
```

## Run Locally

```bash
git clone --recurse-submodules https://github.com/Holius/ASTeroidMermaid.git
cd ASTeroidMermaid
# dir is relative/absolute path to Go code
# module is module name from go.mod beloning to the dir
go run . --dir zarf/src/internal --module github.com/zarf-dev/zarf
# note the above command only processes a subset of "zarf"
# provide the root directory of code to process entire codebase
cat out.mermaid
```

The output can pasted into [mermaid.live](https://mermaid.live/) for quick rendering.

Change `main.go` to point to a different directory with Go Module and change `moduleName` accordingly

## AWK Script

The GNU AWK script generates `expected_go_metadata.go` via `make awk`.
This file is used in regression testing.  It is a secondary method of generating metadata
to verify the Go code in different way with the belief that multiple methods of verification
is better than one.

## Tests

The expected outputs can be found in the `expected_*.go` files.  This was done to keep the actual `*_test.go` files lean.

```go
go test .
```