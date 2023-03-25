package options

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGetRootOptions(t *testing.T) {
	opts := GetStartOptions()
	assert.NotNil(t, opts)
}

func TestInitFlags(t *testing.T) {
	cmd := cobra.Command{}
	opts := GetStartOptions()
	opts.InitFlags(&cmd)
}

func TestStartOptions_SetZeroValues(t *testing.T) {
	opts := GetStartOptions()
	opts.SetZeroValues()
}
