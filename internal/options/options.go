package options

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var s3CleanerOptions = &S3CleanerOptions{}

// S3CleanerOptions contains frequent command line and application options.
type S3CleanerOptions struct {
	// AccessKey is the access key credentials for accessing AWS over client
	AccessKey string
	// SecretKey is the secret key credentials for accessing AWS over client
	SecretKey string
	// BucketName is the name of target bucket
	BucketName string
	// Region is the region of the target bucket
	Region string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}

// GetS3CleanerOptions returns the pointer of S3CleanerOptions
func GetS3CleanerOptions() *S3CleanerOptions {
	return s3CleanerOptions
}

func (opts *S3CleanerOptions) SetAccessCredentialsFromEnv(cmd *cobra.Command) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("aws")
	if err := viper.BindEnv("access_key", "secret_key", "bucket_name", "region"); err != nil {
		return err
	}

	if accessKey := viper.Get("access_key"); accessKey != nil {
		opts.AccessKey = fmt.Sprintf("%v", accessKey)
	} else {
		_ = cmd.MarkFlagRequired("accessKey")
	}

	if secretKey := viper.Get("secret_key"); secretKey != nil {
		opts.SecretKey = fmt.Sprintf("%v", secretKey)
	} else {
		_ = cmd.MarkFlagRequired("secretKey")
	}

	if bucketName := viper.Get("bucket_name"); bucketName != nil {
		opts.BucketName = fmt.Sprintf("%v", bucketName)
	} else {
		_ = cmd.MarkFlagRequired("bucketName")
	}

	if region := viper.Get("region"); region != nil {
		opts.Region = fmt.Sprintf("%v", region)
	} else {
		_ = cmd.MarkFlagRequired("region")
	}

	return nil
}
