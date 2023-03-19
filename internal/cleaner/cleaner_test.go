package cleaner

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	options2 "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/stretchr/testify/assert"
)

var (
	listObjectsErr           error
	getObjectsErr            error
	deleteObjectErr          error
	defaultListObjectsOutput = &s3.ListObjectsOutput{
		Name:        aws.String(""),
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
		Contents: []*s3.Object{
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
				Key:          aws.String("file1.txt"),
				StorageClass: aws.String("STANDARD"),
				Size:         aws.Int64(1000),
				LastModified: aws.Time(time.Now()),
			},
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
				Key:          aws.String("file2.txt"),
				StorageClass: aws.String("STANDARD"),
				Size:         aws.Int64(2000),
				LastModified: aws.Time(time.Now()),
			},
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
				Key:          aws.String("file3.txt"),
				StorageClass: aws.String("STANDARD"),
				Size:         aws.Int64(3000),
				LastModified: aws.Time(time.Now()),
			},
		},
	}
	defaultDeleteObjectOutput = &s3.DeleteObjectOutput{
		DeleteMarker:   nil,
		RequestCharged: nil,
		VersionId:      nil,
	}
	mockLogger = logging.GetLogger()
)

type mockS3Client struct {
	s3iface.S3API
}

// ListObjects mocks the S3API ListObjects method
func (m *mockS3Client) ListObjects(obj *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return defaultListObjectsOutput, listObjectsErr
}

// GetObject mocks the S3API GetObject method
func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	bytes, err := os.Open(*input.Key)
	if err != nil {
		return nil, err
	}

	return &s3.GetObjectOutput{
		AcceptRanges:  aws.String("bytes"),
		Body:          bytes,
		ContentLength: aws.Int64(1000),
		ContentType:   aws.String("text/plain"),
		ETag:          aws.String("d73a503d212d9279e6b2ed8ac6bb81f3"),
	}, getObjectsErr
}

func (m *mockS3Client) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return defaultDeleteObjectOutput, deleteObjectErr
}

func TestStartCleaning(t *testing.T) {
	m := &mockS3Client{}

	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.Nil(t, err)

	// reset zero values again
	startOpts.DryRun = false
	startOpts.AutoApprove = false
}

func TestStartCleaningDryRun(t *testing.T) {
	m := &mockS3Client{}

	listObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.Nil(t, err)

	// reset zero values again
	listObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
}

func TestStartCleaningDryRun1(t *testing.T) {
	m := &mockS3Client{}

	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 10
	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.Nil(t, err)

	// reset zero values again
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
}

func TestStartCleaningDryRun2(t *testing.T) {
	m := &mockS3Client{}

	listObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	startOpts.MinFileSizeInMb = 10
	startOpts.MaxFileSizeInMb = 0

	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.Nil(t, err)

	// reset zero values again
	listObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
}

func TestStartCleaningListError(t *testing.T) {
	m := &mockS3Client{}

	listObjectsErr = errors.New("dummy list error")
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.NotNil(t, err)

	// reset zero values again
	listObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
}

func TestStartCleaningDeleteError(t *testing.T) {
	m := &mockS3Client{}

	deleteObjectErr = errors.New("dummy delete error")
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, mockLogger)
	assert.NotNil(t, err)

	// reset zero values again
	deleteObjectErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
}
