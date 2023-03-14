package start

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// StartCmd represents the bar command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger()

		rootOpts := cmd.Context().Value(options.CtxKey{}).(*options.RootOptions)
		fmt.Println(rootOpts)

		//rootOpts := root.GetRootOptions()
		sess, err := aws.CreateSession(rootOpts)
		if err != nil {
			logger.Error("an error occured while creating session", zap.Error(err))
			return err
		}

		// obtain S3 client with initialized session
		svc := s3.New(sess)

		logger.Debug("trying to find files on bucket", zap.String("bucketName", rootOpts.BucketName),
			zap.String("region", rootOpts.Region))

		allFiles, err := aws.GetAllFiles(svc, rootOpts)
		if err != nil {
			return err
		}

		for _, v := range allFiles.Contents {
			logger.Info(*v.Key)
			logger.Info(fmt.Sprintf("%v", *v.Size))
			logger.Info(fmt.Sprintf("%v", *v.LastModified))
			logger.Info(fmt.Sprintf("%v", v.LastModified.Add(120*time.Hour).Before(time.Now())))
		}

		return nil
	},
}
