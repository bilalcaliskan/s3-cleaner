package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	res := Contains([]string{"size", "lastModificationDate"}, "size")
	assert.True(t, res)
}

func TestNotContains(t *testing.T) {
	res := Contains([]string{"size", "lastModificationDate"}, "sizee")
	assert.False(t, res)
}
