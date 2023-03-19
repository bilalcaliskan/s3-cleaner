package cleaner

import (
	"bytes"

	"github.com/aws/aws-sdk-go/service/s3"
	root "github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	start "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
	"github.com/rs/zerolog"
)

func StartCleaning(rootOpts *root.RootOptions, startOpts *start.StartOptions, logger zerolog.Logger) error {
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
