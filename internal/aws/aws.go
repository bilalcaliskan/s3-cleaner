package aws

import (
	"log"

	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/rs/zerolog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
)

var logger zerolog.Logger

func init() {
	logger = logging.GetLogger()
}

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
		Prefix: aws.String(opts.FileNamePrefix),
	})
	if err != nil {
		return res, err
	}

	//for _, v := range res.Contents {
	//	logger.Info(*v.Key)
	//}

	return res, nil
}

func DeleteFiles(svc s3iface.S3API, bucketName string, slice []*s3.Object, dryRun bool) error {
	for _, v := range slice {
		logger.Info().Str("key", *v.Key).Time("lastModifiedDate", *v.LastModified).
			Float64("size", float64(*v.Size)/1000000).Msg("will try to delete file")
		//logger.Debug(fmt.Sprintf("will try to delete file %s with last modification date %v and size %f MB", *v.Key, *v.LastModified, float64(*v.Size)/1000000))

		if !dryRun {
			_, err := svc.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(*v.Key),
			})

			if err != nil {
				return err
			}

			log.Printf("successfully deleted file %s", *v.Key)
			logger.Info().Str("key", *v.Key).Msg("successfully deleted file")
		}
	}

	return nil
}
