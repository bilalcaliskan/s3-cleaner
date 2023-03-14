package root

import (
	"context"
	"os"

	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/cmd/start"
	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/version"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	opts   *options.RootOptions
	ver    = version.Get()
)

func init() {
	logger = logging.GetLogger()
	opts = options.GetRootOptions()
	options.InitFlags(rootCmd, opts)

	rootCmd.AddCommand(start.StartCmd)
}

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "s3-cleaner",
	Short:   "This tool finds the desired files in a bucket and cleans them",
	Long:    ``,
	Version: ver.GitVersion,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		//bannerBytes, _ := os.ReadFile("banner.txt")
		//banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))

		if opts.VerboseLog {
			logging.Atomic.SetLevel(zap.DebugLevel)
		}

		logger.Info("s3-cleaner is started",
			zap.String("appVersion", ver.GitVersion),
			zap.String("goVersion", ver.GoVersion),
			zap.String("goOS", ver.GoOs),
			zap.String("goArch", ver.GoArch),
			zap.String("gitCommit", ver.GitCommit),
			zap.String("buildDate", ver.BuildDate))

		ctx := context.WithValue(context.Background(), options.CtxKey{}, opts)
		cmd.SetContext(ctx)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
