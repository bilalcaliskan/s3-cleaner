package aws

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/stretchr/testify/assert"
)

func TestGetAllFilesHappyPath(t *testing.T) {
	m := &MockS3Client{}
	DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
		},
	}

	result, err := GetAllFiles(m, options.GetRootOptions())
	assert.NotEmpty(t, result)
	assert.Nil(t, err)
}

func TestGetAllFilesFailedListObjectsCall(t *testing.T) {
	m := &MockS3Client{}
	ListObjectsErr = errors.New("dummy error thrown")
	_, err := GetAllFiles(m, options.GetRootOptions())
	assert.NotNil(t, err)
	ListObjectsErr = nil
}

func TestDeleteFilesHappyPath(t *testing.T) {
	var input []*s3.Object
	m := &MockS3Client{}
	DeleteObjectErr = nil

	err := DeleteFiles(m, "dummy bucket", input, false, MockLogger)
	assert.Nil(t, err)
}

func TestDeleteFilesHappyPathDryRun(t *testing.T) {
	var input []*s3.Object
	m := &MockS3Client{}
	DeleteObjectErr = nil

	err := DeleteFiles(m, "dummy bucket", input, true, MockLogger)
	assert.Nil(t, err)
}

func TestDeleteFilesFailedDeleteObjectCall(t *testing.T) {
	var input []*s3.Object
	for i := 0; i < 3; i++ {
		o := s3.Object{Key: aws.String("hello-world"), LastModified: aws.Time(time.Now()), Size: aws.Int64(10000000)}
		input = append(input, &o)
	}

	m := &MockS3Client{}
	DeleteObjectErr = errors.New("dummy error")
	err := DeleteFiles(m, "dummy bucket", input, false, MockLogger)
	assert.NotNil(t, err)
	DeleteObjectErr = nil
}

func TestCreateSession(t *testing.T) {
	sess, err := CreateSession(options.GetRootOptions())
	assert.Nil(t, err)
	assert.NotNil(t, sess)
}
