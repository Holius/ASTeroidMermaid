package main

import (
	"cmp"
	"slices"
	"testing"

	"github.com/holius/asteroid_mermaid/filedata"
	"github.com/stretchr/testify/assert"
)

func TestGoMetadataFromDirectory(t *testing.T) {
	// Act
	gm, err := filedata.GoMetadataFromDirectory("zarf")

	// Assert
	assert.NoError(t, err)
	slices.SortFunc(gm, func(a, b filedata.GoMetadata) int {
		return cmp.Compare(a.File, b.File)
	})
	slices.SortFunc(filedata.GoMetadataExpected, func(a, b filedata.GoMetadata) int {
		return cmp.Compare(a.File, b.File)
	})
	for i := range gm {
		assert.Equal(t, filedata.GoMetadataExpected[i], gm[i])
		//if i == 1 {
		//	break
		//}
	}
}
