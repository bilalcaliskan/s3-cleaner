package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
)

//var logger *zap.Logger
//
//func init() {
//	logger = logging.GetLogger()
//}

// CreateSession initializes session with provided credentials
func CreateSession(opts *options.RootOptions) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})

	return sess, err
}

func GetAllFiles(svc s3iface.S3API, opts *options.RootOptions) (*s3.ListObjectsOutput, error) {
	var err error
	var res *s3.ListObjectsOutput

	// fetch all the objects in target bucket
	res, err = svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(opts.BucketName),
	})
	if err != nil {
		return res, err
	}

	//for _, v := range res.Contents {
	//	logger.Info(*v.Key)
	//}

	return res, nil
}
