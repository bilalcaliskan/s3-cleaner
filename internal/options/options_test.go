package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetS3CleanerOptions function tests if GetS3CleanerOptions function running properly
func TestGetS3CleanerOptions(t *testing.T) {
	t.Log("fetching default options.S3Cleaner")
	opts := GetS3CleanerOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.S3CleanerOptions, %v\n", opts)
}
