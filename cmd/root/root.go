package root

import (
	"context"
	"log"
	"os"
	"strings"

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
	options.InitFlags(rootCmd, opts)

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

		log.Printf("s3-cleaner is started appVersion=%s goVersion=%s goOS=%s goArch=%s gitCommit=%s buildDate=%s",
			ver.GitVersion, ver.GoVersion, ver.GoOs, ver.GoArch, ver.GitCommit, ver.BuildDate)

		cmd.SetContext(context.WithValue(context.Background(), options.CtxKey{}, opts))

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
