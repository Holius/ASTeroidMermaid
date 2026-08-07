package packagedata

import (
	"maps"
	"path/filepath"
	"slices"
	"testing"

	"github.com/holius/asteroid_mermaid/filedata"
	"github.com/stretchr/testify/assert"
)

// func GoPackageDataFromGoMetadata(gmSlice []GoMetadata) (map[string]*GoPackagaedata, error) {
func TestGoPackageDataFromGoMetadata(t *testing.T) {

	// Arrange
	input := []filedata.GoMetadata{}
	packages := []string{"zarf", "zarf/src/internal/template", "zarf/src/internal/dns", "zarf/src/internal/api/v1alpha1", "zarf/src/internal/pkgcfg", "zarf/src/internal/git"}
	for _, fd := range filedata.GoFileDataExpected {
		if slices.Contains(packages, filepath.Dir(fd.File)) {
			input = append(input, fd)
		}
	}

	// Act
	pg, err := GoPackageDataFromGoFileData(input, true)

	// Assert
	assert.NoError(t, err)

	// make sure keys are the same
	expectedKeys := slices.Collect(maps.Keys(GoPackageDataExpected))
	actualKeys := slices.Collect(maps.Keys(pg))
	slices.Sort(expectedKeys)
	slices.Sort(actualKeys)
	assert.Equal(t, expectedKeys, actualKeys)

	// makes sure values are the same
	for k := range GoPackageDataExpected {
		slices.Sort(GoPackageDataExpected[k].Imports)
		slices.Sort(pg[k].Imports)
		assert.Equal(t, GoPackageDataExpected[k], pg[k])
	}
}

func TestGoPackageDataFromGoMetadataWithoutTests(t *testing.T) {

	// Arrange
	input := []filedata.GoMetadata{}
	packages := []string{"zarf", "zarf/src/internal/template", "zarf/src/internal/dns", "zarf/src/internal/api/v1alpha1", "zarf/src/internal/pkgcfg", "zarf/src/internal/git"}
	for _, fd := range filedata.GoFileDataExpected {
		if slices.Contains(packages, filepath.Dir(fd.File)) {
			input = append(input, fd)
		}
	}

	// Act
	pg, err := GoPackageDataFromGoFileData(input, false)

	// Assert
	assert.NoError(t, err)

	// make sure keys are the same
	expectedKeys := slices.Collect(maps.Keys(goPackageDataExpectedWithoutTests))
	actualKeys := slices.Collect(maps.Keys(pg))
	slices.Sort(expectedKeys)
	slices.Sort(actualKeys)
	assert.Equal(t, expectedKeys, actualKeys)

	// makes sure values are the same
	for k := range goPackageDataExpectedWithoutTests {
		slices.Sort(goPackageDataExpectedWithoutTests[k].Imports)
		slices.Sort(pg[k].Imports)
		assert.Equal(t, goPackageDataExpectedWithoutTests[k], pg[k])
	}
}
