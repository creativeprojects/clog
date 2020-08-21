package clog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultOutput(t *testing.T) {
	assert.Nil(t, stdout)
	assert.Nil(t, stderr)
	assert.NotNil(t, Stdout())
	assert.NotNil(t, Stderr())
	assert.NotNil(t, stdout)
	assert.NotNil(t, stderr)
}
