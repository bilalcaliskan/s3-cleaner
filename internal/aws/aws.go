package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/options"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

// CreateSession initializes session with provided credentials
func CreateSession(opts *options.S3CleanerOptions) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})

	return sess, err
}

func GetAllFiles(svc s3iface.S3API, opts *options.S3CleanerOptions) error {
	var err error

	// fetch all the objects in target bucket
	listResult, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(opts.BucketName),
	})
	if err != nil {
		return err
	}

	for _, v := range listResult.Contents {
		logger.Info("found file", zap.String("name", *v.Key))
		logger.Info(v.String())
	}

	return nil
}
