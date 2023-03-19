package start

import (
	"fmt"

	rootopts "github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/cleaner"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/utils"
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
			return cleaner.StartCleaning(cmd.Context().Value(rootopts.OptsKey{}).(*rootopts.RootOptions), startOpts, logger)
		},
	}
)
