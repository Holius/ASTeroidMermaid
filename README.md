# ASTeroid Mermaid

This app maps public/exported functions in a package
to other intra packages that it imports
- "intra package" means a package that exists within the same module

Below is expected output from subset of Zarf source code; this shows that `v1alpha1` imports `v1beta1` (and showing all Public functions):
Using GitHub sss7526's [resistor](https://github.com/sss7526/resistor) project as an example of what an ASTeroid Mermaid generation looks like.  (Yes, the name was made up while very tired.)
```mermaid
graph TD
  subgraph resistor
  AnalyzeResistor
  BandRole.String
  BandRolesForCount
  ValidColorsForRole
  DecodeBands
  EncodeBands
  EncodeBandsSimple
  Color.String
  BodyColors
  ESeries.String
  ParseESeries
  AllESeries
  RoundingMode.String
  ParseRoundingMode
  AllRoundingModes
  PackageType.String
  ParsePackageType
  AllPackageTypes
  DigitColors
  MultiplierColors
  ToleranceColors
  TempCoeffColors
  NearestStandard
  InferResistor
  SelectStandardResistor
  DecodeSMD
  EncodeSMD
  end
  subgraph resistor/cmd/resistor-cli
  end
  resistor/cmd/resistor-cli --> resistor/cmd/resistor-cli/cmd
  subgraph resistor/cmd/resistor-cli/cmd
  Execute
  end
  resistor/cmd/resistor-cli/cmd --> resistor
  resistor/cmd/resistor-cli/cmd --> resistor/internal/cli
  subgraph resistor/cmd/resistor-server
  noDirFS.Open
  end
  resistor/cmd/resistor-server --> resistor/web
  subgraph resistor/cmd/resistor-tui
  end
  resistor/cmd/resistor-tui --> resistor/cmd/resistor-tui/app
  subgraph resistor/cmd/resistor-tui/app
  NewAnalyzeView
  AnalyzeView.Resize
  AnalyzeView.Init
  AnalyzeView.Update
  AnalyzeView.View
  BaseView.Resize
  NewInferView
  InferView.Resize
  InferView.Init
  InferView.Update
  InferView.View
  menuItem.Title
  menuItem.Description
  menuItem.FilterValue
  NewMenu
  MenuView.Resize
  MenuView.Init
  MenuView.Update
  MenuView.View
  New
  AppModel.Init
  AppModel.Update
  AppModel.View
  NewPlaceholderView
  PlaceholderView.Init
  PlaceholderView.Update
  PlaceholderView.View
  NewSelectView
  SelectView.Resize
  SelectView.Init
  SelectView.Update
  SelectView.View
  NewSMDView
  SMDView.Resize
  SMDView.Init
  SMDView.Update
  SMDView.View
  end
  resistor/cmd/resistor-tui/app --> resistor
  subgraph resistor/cmd/resistor-wasm
  end
  resistor/cmd/resistor-wasm --> resistor
  subgraph resistor/internal/cli
  PrintHeader
  PrintBands
  OutputJSONSuccess
  OutputJSONError
  Respond
  end
  resistor/internal/cli --> resistor
  subgraph resistor/web
  end
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