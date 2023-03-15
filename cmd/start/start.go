package start

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/service/s3"
	rootopts "github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	logger = logging.GetLogger()
	startOpts = options.GetStartOptions()

	options.InitFlags(StartCmd, startOpts)
}

var (
	logger          *zap.Logger
	ValidSortByOpts = []string{"size", "lastModificationDate"}
	startOpts       *options.StartOptions
	// StartCmd represents the bar command
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start subcommand starts the app, finds and clears desired files",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if startOpts.MinFileSizeInMb > startOpts.MaxFileSizeInMb {
				return fmt.Errorf("minFileSizeInMb should be lower than maxFileSizeInMb")
			}

			if !utils.Contains(ValidSortByOpts, startOpts.SortBy) {
				return fmt.Errorf("no such sortBy option called %s, valid options are %v", startOpts.SortBy,
					ValidSortByOpts)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			rootOpts := cmd.Context().Value(rootopts.CtxKey{}).(*rootopts.RootOptions)

			sess, err := aws.CreateSession(rootOpts)
			if err != nil {
				logger.Error("an error occurred while creating session", zap.Error(err))
				return err
			}

			svc := s3.New(sess)

			logger.Debug("trying to find files on bucket", zap.String("bucketName", rootOpts.BucketName),
				zap.String("region", rootOpts.Region))

			allFiles, err := aws.GetAllFiles(svc, rootOpts)
			if err != nil {
				return err
			}

			var res []*s3.Object
			for _, v := range allFiles.Contents {
				if *v.Size > startOpts.MinFileSizeInMb*1000000 && *v.Size < startOpts.MaxFileSizeInMb*1000000 {
					res = append(res, v)
				}
			}

			switch startOpts.SortBy {
			case "lastModificationDate":
				sort.Slice(res, func(i, j int) bool {
					return res[i].LastModified.Before(*res[j].LastModified)
				})
			case "size":
				sort.Slice(res, func(i, j int) bool {
					return *res[i].Size < *res[j].Size
				})
			}

			logger.Debug(fmt.Sprintf("length of result slice is %d", len(res)))
			return aws.DeleteFiles(svc, rootOpts.BucketName, res[:len(res)-startOpts.KeepLastNFiles], startOpts.DryRun)
		},
	}
)
