package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func GoPackageDataFromGoMetadata(gmSlice []GoMetadata) (map[string]*GoPackagaedata, error) {
func TestGoPackageDataFromGoMetadata(t *testing.T) {

	// Act
	pg, err := GoPackageDataFromGoMetadata(goMetadataExpected[:11], true)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, goPackageDataExpected, pg)
}

func TestGoPackageDataFromGoMetadataWithoutTests(t *testing.T) {

	// Act
	pg, err := GoPackageDataFromGoMetadata(goMetadataExpected[:11], false)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, goPackageDataExpectedWithoutTests, pg)
}
