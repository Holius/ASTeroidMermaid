package main

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoMetadataFromDirectory(t *testing.T) {
	// Act
	gm, err := GoMetadataFromDirectory("zarf")

	// Assert
	assert.NoError(t, err)
	slices.SortFunc(gm, func(a, b GoMetadata) int {
		return cmp.Compare(a.File, b.File)
	})
	slices.SortFunc(goMetadataExpected, func(a, b GoMetadata) int {
		return cmp.Compare(a.File, b.File)
	})
	assert.Equal(t, goMetadataExpected, gm)
}
