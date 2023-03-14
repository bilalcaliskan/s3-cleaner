package options

import "github.com/spf13/cobra"

var startOptions = &StartOptions{}

// StartOptions contains frequent command line and application options.
type StartOptions struct {
	MinFileSizeInBytes int64
	MaxFileSizeInBytes int64
	FileExtensions     string
	KeepLastNFiles     int
	DryRun             bool
}

// GetStartOptions returns the pointer of S3CleanerOptions
func GetStartOptions() *StartOptions {
	return startOptions
}

func InitFlags(cmd *cobra.Command, opts *StartOptions) {
	cmd.Flags().Int64VarP(&opts.MinFileSizeInBytes, "minFileSizeInBytes", "", 10000000, "")
	cmd.Flags().Int64VarP(&opts.MaxFileSizeInBytes, "maxFileSizeInBytes", "", 15000000, "")
	cmd.Flags().StringVarP(&opts.FileExtensions, "fileExtensions", "", "", "")
	cmd.Flags().IntVarP(&opts.KeepLastNFiles, "keepLastNFiles", "", 1, "")
	cmd.Flags().BoolVarP(&opts.DryRun, "dryRun", "", false, "")
}
