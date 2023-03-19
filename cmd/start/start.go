package start

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	rootopts "github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func init() {
	logger = logging.GetLogger()
	startOpts = options.GetStartOptions()

	options.InitFlags(StartCmd, startOpts)
}

var (
	logger          zerolog.Logger
	ValidSortByOpts = []string{"size", "lastModificationDate"}
	startOpts       *options.StartOptions
	// StartCmd represents the bar command
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "start subcommand starts the app, finds and clears desired files",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			logger = cmd.Context().Value(rootopts.LoggerKey{}).(zerolog.Logger)

			if startOpts.MinFileSizeInMb > startOpts.MaxFileSizeInMb && (startOpts.MinFileSizeInMb != 0 && startOpts.MaxFileSizeInMb != 0) {
				err := fmt.Errorf("minFileSizeInMb should be lower than maxFileSizeInMb")
				logger.Error().Str("error", err.Error()).Msg("an error occured while validating flags")
				return err
			}

			if !utils.Contains(ValidSortByOpts, startOpts.SortBy) {
				err := fmt.Errorf("no such sortBy option called %s, valid options are %v", startOpts.SortBy,
					ValidSortByOpts)
				logger.Error().Str("error", err.Error()).Msg("an error occurred while validating flags")
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			rootOpts := cmd.Context().Value(rootopts.OptsKey{}).(*rootopts.RootOptions)

			sess, err := aws.CreateSession(rootOpts)
			if err != nil {
				logger.Error().Str("error", err.Error()).Msg("an error occurred while creating session")
				return err
			}

			svc := s3.New(sess)
			logger.Info().Str("bucket", rootOpts.BucketName).Str("region", rootOpts.Region).Msg("trying " +
				"to find files on target bucket")

			allFiles, err := aws.GetAllFiles(svc, rootOpts)
			if err != nil {
				return err
			}

			var res []*s3.Object
			for _, v := range allFiles.Contents {
				if strings.HasSuffix(*v.Key, "/") {
					logger.Debug().Str("key", *v.Key).Msg("object has directory suffix, skipping that one")
					continue
				}

				if (startOpts.MinFileSizeInMb == 0 && startOpts.MaxFileSizeInMb != 0) && *v.Size < startOpts.MaxFileSizeInMb*1000000 { // case 2
					res = append(res, v)
				} else if (startOpts.MinFileSizeInMb != 0 && startOpts.MaxFileSizeInMb == 0) && *v.Size >= startOpts.MinFileSizeInMb*1000000 { // case 3
					res = append(res, v)
				} else if startOpts.MinFileSizeInMb == 0 && startOpts.MaxFileSizeInMb == 0 { // case 1
					res = append(res, v)
				} else if startOpts.MinFileSizeInMb != 0 && startOpts.MaxFileSizeInMb != 0 && (*v.Size >= startOpts.MinFileSizeInMb*1000000 && *v.Size < startOpts.MaxFileSizeInMb*1000000) { // case 4
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

			if len(res) < startOpts.KeepLastNFiles {
				logger.Warn().Str("bucket", rootOpts.BucketName).Str("region", rootOpts.Region).Msg("no " +
					"file found with specified specs on target bucket")
				return nil
			}

			keys := utils.GetKeysOnly(res)
			var buffer bytes.Buffer
			for _, v := range keys {
				buffer.WriteString(v)
			}

			if !startOpts.AutoApprove {
				logger.Info().Any("files", keys).Msg("these files will be removed if you approve:")

				prompt := promptui.Prompt{
					Label:     "Delete Files?",
					IsConfirm: true,
					Validate: func(s string) error {
						if len(s) == 1 {
							return nil
						}

						return errors.New("invalid input")
					},
				}

				if _, err := prompt.Run(); err != nil {
					return err
				}
			}

			targetObjects := res[:len(res)-startOpts.KeepLastNFiles]
			if len(targetObjects) == 0 {
				logger.Info().Str("bucket", rootOpts.BucketName).Str("region", rootOpts.Region).
					Msg("no deletable file found on the target bucket")
				return nil
			}

			logger.Info().Any("files", keys).Msg("trying to delete files")

			if err := aws.DeleteFiles(svc, rootOpts.BucketName, targetObjects, startOpts.DryRun, logger); err != nil {
				logger.Error().Str("error", err.Error()).Msg("an error occurred while deleting target files")
				return err
			}

			if startOpts.DryRun {
				logger.Info().Msg("skipping object deletion since --dryRun flag is passed")
			}

			return nil
		},
	}
)
