package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoPackagedataToMermaid(t *testing.T) {

	// Act
	mermaid := GoPackagedataToMermaid(mermaidInput, "github.com/zarf-dev/zarf")

	// Assert
	assert.Equal(t, mermaidExpected, mermaid)

}
