package glib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, AsciiType(16), ASCII_GRAPH)
}
