package packagedata

import (
	"testing"

	"github.com/holius/asteroid_mermaid/filedata"
	"github.com/stretchr/testify/assert"
)

// func GoPackageDataFromGoMetadata(gmSlice []GoMetadata) (map[string]*GoPackagaedata, error) {
func TestGoPackageDataFromGoMetadata(t *testing.T) {

	// Act
	pg, err := GoPackageDataFromGoMetadata(filedata.GoMetadataExpected[:11], true)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, GoPackageDataExpected, pg)
}

func TestGoPackageDataFromGoMetadataWithoutTests(t *testing.T) {

	// Act
	pg, err := GoPackageDataFromGoMetadata(filedata.GoMetadataExpected[:11], false)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, goPackageDataExpectedWithoutTests, pg)
}
