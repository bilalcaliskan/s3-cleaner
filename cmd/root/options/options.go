package options

import "github.com/spf13/cobra"

var rootOptions = &RootOptions{}

type (
	OptsKey   struct{}
	LoggerKey struct{}
)

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	// AccessKey is the access key credentials for accessing AWS over client
	AccessKey string
	// SecretKey is the secret key credentials for accessing AWS over client
	SecretKey string
	// BucketName is the name of target bucket
	BucketName string
	// FileNamePrefix is the prefix of target bucket objects, means it can be used for folder-based object grouping buckets
	FileNamePrefix string
	// Region is the region of the target bucket
	Region string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}

// GetRootOptions returns the pointer of S3CleanerOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func InitFlags(cmd *cobra.Command, opts *RootOptions) {
	cmd.PersistentFlags().StringVarP(&opts.BucketName, "bucketName", "", "", "name of "+
		"the target bucket on S3 (default \"\")")
	cmd.PersistentFlags().StringVarP(&opts.FileNamePrefix, "fileNamePrefix", "", "",
		"folder name of target bucket objects, means it can be used for folder-based object grouping buckets (default \"\")")
	cmd.PersistentFlags().StringVarP(&opts.AccessKey, "accessKey", "", "",
		"access key credential to access S3 bucket (default \"\")")
	cmd.PersistentFlags().StringVarP(&opts.SecretKey, "secretKey", "", "",
		"secret key credential to access S3 bucket (default \"\")")
	cmd.PersistentFlags().StringVarP(&opts.Region, "region", "", "",
		"region of the target bucket on S3 (default \"\")")
	cmd.PersistentFlags().BoolVarP(&opts.VerboseLog, "verbose", "", false,
		"verbose output of the logging library (default false)")

	_ = cmd.MarkFlagRequired("bucketName")
	_ = cmd.MarkFlagRequired("accessKey")
	_ = cmd.MarkFlagRequired("secretKey")
	_ = cmd.MarkFlagRequired("region")
}
