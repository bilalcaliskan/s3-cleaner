package options

var startOptions = &StartOptions{}

// StartOptions contains frequent command line and application options.
type StartOptions struct {
	MinSizeInBytes int
	MaxSizeInBytes int
	FileExtensions string
	KeepLastNFiles int
}

// GetStartOptions returns the pointer of S3CleanerOptions
func GetStartOptions() *StartOptions {
	return startOptions
}
