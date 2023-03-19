package cleaner

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	root "github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	start "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
	"github.com/rs/zerolog"
)

func StartCleaning(svc s3iface.S3API, rootOpts *root.RootOptions, startOpts *start.StartOptions, logger zerolog.Logger) error {
	allFiles, err := aws.GetAllFiles(svc, rootOpts)
	if err != nil {
		return err
	}

	res := getProperObjects(startOpts, allFiles, logger)
	sortObjects(res, startOpts)

	targetObjects := res[:len(res)-startOpts.KeepLastNFiles]
	if err := checkLength(targetObjects); err != nil {
		logger.Warn().Str("bucket", rootOpts.BucketName).Str("region", rootOpts.Region).
			Msg(err.Error())
		return nil
	}

	keys := utils.GetKeysOnly(targetObjects)
	var buffer bytes.Buffer
	for _, v := range keys {
		buffer.WriteString(v)
	}

	if startOpts.DryRun {
		logger.Info().Msg("skipping object deletion since --dryRun flag is passed")
		return nil
	}

	if err := promptDeletion(startOpts, logger, keys); err != nil {
		logger.Error().Str("error", err.Error()).Msg("an error occurred while prompting file deletion")
		return err
	}

	logger.Info().Any("files", keys).Msg("will attempt to delete these files")
	if err := aws.DeleteFiles(svc, rootOpts.BucketName, targetObjects, startOpts.DryRun, logger); err != nil {
		logger.Error().Str("error", err.Error()).Msg("an error occurred while deleting target files")
		return err
	}

	return nil
}
