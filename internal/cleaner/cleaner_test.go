package cleaner

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	options2 "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	internalAws "github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/stretchr/testify/assert"
)

func TestStartCleaning(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(1000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(2000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(3000),
			LastModified: aws.Time(time.Now()),
		},
	}

	internalAws.ListObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.Nil(t, err)
}

func TestStartCleaningDryRun(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(1000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(2000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(3000),
			LastModified: aws.Time(time.Now()),
		},
	}

	internalAws.ListObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.Nil(t, err)

	// reset zero values again
	internalAws.ListObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	internalAws.DefaultListObjectsOutput = &s3.ListObjectsOutput{
		Name:        aws.String(""),
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
	}
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
}

func TestStartCleaningDryRun1(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(1000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(2000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(3000),
			LastModified: aws.Time(time.Now()),
		},
	}

	internalAws.ListObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 10
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.Nil(t, err)

	// reset zero values again
	internalAws.ListObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	internalAws.DefaultListObjectsOutput = &s3.ListObjectsOutput{
		Name:        aws.String(""),
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
	}
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
}

func TestStartCleaningDryRun2(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(1000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(2000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(3000),
			LastModified: aws.Time(time.Now()),
		},
	}

	internalAws.ListObjectsErr = nil
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = true
	startOpts.AutoApprove = true
	startOpts.MinFileSizeInMb = 10
	startOpts.MaxFileSizeInMb = 0
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.Nil(t, err)

	// reset zero values again
	internalAws.ListObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	internalAws.DefaultListObjectsOutput = &s3.ListObjectsOutput{
		Name:        aws.String(""),
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
	}
	startOpts.MinFileSizeInMb = 0
	startOpts.MaxFileSizeInMb = 0
}

func TestStartCleaningListError(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("../../mock/file1.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(1000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
			Key:          aws.String("../../mock/file2.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(2000),
			LastModified: aws.Time(time.Now()),
		},
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
			Key:          aws.String("../../mock/file3.txt"),
			StorageClass: aws.String("STANDARD"),
			Size:         aws.Int64(3000),
			LastModified: aws.Time(time.Now()),
		},
	}

	internalAws.ListObjectsErr = errors.New("dummy list error")
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.NotNil(t, err)

	// reset zero values again
	internalAws.ListObjectsErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
	internalAws.DefaultListObjectsOutput = &s3.ListObjectsOutput{
		Name:        aws.String(""),
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
	}
}

func TestStartCleaningDeleteError(t *testing.T) {
	m := &internalAws.MockS3Client{}
	internalAws.DefaultListObjectsOutput.Contents = []*s3.Object{
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
	}

	internalAws.DeleteObjectErr = errors.New("dummy delete error")
	startOpts := options2.GetStartOptions()
	startOpts.DryRun = false
	startOpts.AutoApprove = true
	err := StartCleaning(m, options.GetRootOptions(), startOpts, internalAws.MockLogger)
	assert.NotNil(t, err)

	// reset zero values again
	internalAws.DeleteObjectErr = nil
	startOpts.DryRun = false
	startOpts.AutoApprove = false
}
