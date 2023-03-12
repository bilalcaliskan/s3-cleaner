package options

var s3CleanerOptions = &S3CleanerOptions{}

// S3CleanerOptions contains frequent command line and application options.
type S3CleanerOptions struct {
	// Foo is the dummy option
	Foo string
}

// GetS3CleanerOptions returns the pointer of S3CleanerOptions
func GetS3CleanerOptions() *S3CleanerOptions {
	return s3CleanerOptions
}
