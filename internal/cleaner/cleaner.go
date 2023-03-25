package cleaner

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	start "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
	"github.com/rs/zerolog"
)

func StartCleaning(svc s3iface.S3API, startOpts *start.StartOptions, logger zerolog.Logger) error {
	allFiles, err := aws.GetAllFiles(svc, startOpts.RootOptions)
	if err != nil {
		return err
	}

	res := getProperObjects(startOpts, allFiles, logger)
	sortObjects(res, startOpts)

	targetObjects := res[:len(res)-startOpts.KeepLastNFiles]
	if err := checkLength(targetObjects); err != nil {
		logger.Warn().Str("bucket", startOpts.RootOptions.BucketName).Str("region", startOpts.RootOptions.Region).
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
		logger.Warn().Str("error", err.Error()).Msg("an error occurred while prompting file deletion")
		return err
	}

	logger.Info().Any("files", keys).Msg("will attempt to delete these files")
	if err := aws.DeleteFiles(svc, startOpts.RootOptions.BucketName, targetObjects, startOpts.DryRun, logger); err != nil {
		logger.Error().Str("error", err.Error()).Msg("an error occurred while deleting target files")
		return err
	}

	return nil
}
