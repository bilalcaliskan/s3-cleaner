package options

import (
	"github.com/spf13/cobra"
)

var startOptions = &StartOptions{}

// StartOptions contains frequent command line and application options.
type StartOptions struct {
	MinFileSizeInMb int64
	MaxFileSizeInMb int64
	FileExtensions  string
	KeepLastNFiles  int
	DryRun          bool
	SortBy          string
}

// GetStartOptions returns the pointer of S3CleanerOptions
func GetStartOptions() *StartOptions {
	return startOptions
}

func InitFlags(cmd *cobra.Command, opts *StartOptions) {
	cmd.Flags().Int64VarP(&opts.MinFileSizeInMb, "minFileSizeInMb", "", 10,
		"minimum size in mb to clean from target bucket, 0 means no lower limit")
	cmd.Flags().Int64VarP(&opts.MaxFileSizeInMb, "maxFileSizeInMb", "", 15,
		"maximum size in mb to clean from target bucket, 0 means no upper limit")
	cmd.Flags().StringVarP(&opts.FileExtensions, "fileExtensions", "", "",
		"selects the files with defined extensions to clean from target bucket, \"\" means all files (default \"\")")
	cmd.Flags().IntVarP(&opts.KeepLastNFiles, "keepLastNFiles", "", 1,
		"defines how many of the files to skip deletion in specified criteria, 0 means clean them all")
	cmd.Flags().StringVarP(&opts.SortBy, "sortBy", "", "lastModificationDate",
		"defines the ascending order in the specified criteria, valid options are \"lastModificationDate\" and \"size\"")
	cmd.Flags().BoolVarP(&opts.DryRun, "dryRun", "", false, "specifies that if you "+
		"just want to see what to delete or completely delete them all (default false)")
}
