package root

import (
	"context"
	"os"
	"strings"

	"github.com/bilalcaliskan/s3-cleaner/internal/logging"

	"github.com/dimiro1/banner"

	"github.com/bilalcaliskan/s3-cleaner/cmd/root/options"
	"github.com/bilalcaliskan/s3-cleaner/cmd/start"
	"github.com/bilalcaliskan/s3-cleaner/internal/version"
	"github.com/spf13/cobra"
)

var (
	opts *options.RootOptions
	ver  = version.Get()
)

func init() {
	opts = options.GetRootOptions()
	opts.InitFlags(rootCmd)
	if err := opts.SetAccessCredentialsFromEnv(rootCmd); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(start.StartCmd)
}

// Cmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "s3-cleaner",
	Short:   "This tool finds the desired files in a bucket and cleans them",
	Long:    ``,
	Version: ver.GitVersion,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("build/ci/banner.txt"); err == nil {
			bannerBytes, _ := os.ReadFile("build/ci/banner.txt")
			banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
		}

		if opts.VerboseLog {
			logging.EnableDebugLogging()
		}

		logger := logging.GetLogger()
		logger.Info().Str("appVersion", ver.GitVersion).Str("goVersion", ver.GoVersion).Str("goOS", ver.GoOs).
			Str("goArch", ver.GoArch).Str("gitCommit", ver.GitCommit).Str("buildDate", ver.BuildDate).
			Msg("s3-cleaner is started!")

		cmd.SetContext(context.WithValue(cmd.Context(), options.LoggerKey{}, logger))
		cmd.SetContext(context.WithValue(cmd.Context(), options.OptsKey{}, opts))

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
