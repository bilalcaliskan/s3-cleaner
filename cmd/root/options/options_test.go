package options

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGetRootOptions(t *testing.T) {
	opts := GetRootOptions()
	assert.NotNil(t, opts)
}

func TestRootOptions_InitFlags(t *testing.T) {
	cmd := cobra.Command{}
	opts := GetRootOptions()
	opts.InitFlags(&cmd)
}

func TestRootOptions_SetAccessCredentialsFromEnv(t *testing.T) {
	cmd := cobra.Command{}
	opts := GetRootOptions()
	err := opts.SetAccessCredentialsFromEnv(&cmd)
	assert.Nil(t, err)
}
