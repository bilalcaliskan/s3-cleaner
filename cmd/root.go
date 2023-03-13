package cmd

import (
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-cleaner/internal/aws"

	"github.com/bilalcaliskan/s3-cleaner/internal/logging"
	"github.com/bilalcaliskan/s3-cleaner/internal/options"
	"github.com/bilalcaliskan/s3-cleaner/internal/version"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	opts   *options.S3CleanerOptions
	ver    = version.Get()
)

func init() {
	logger = logging.GetLogger()
	opts = options.GetS3CleanerOptions()
	rootCmd.Flags().StringVarP(&opts.BucketName, "bucketName", "", "",
		"name of the target bucket on S3")
	rootCmd.Flags().StringVarP(&opts.AccessKey, "accessKey", "", "",
		"access key credential to access S3 bucket")
	rootCmd.Flags().StringVarP(&opts.SecretKey, "secretKey", "", "",
		"secret key credential to access S3 bucket")
	rootCmd.Flags().StringVarP(&opts.Region, "region", "", "",
		"region of the target bucket on S3")

	if err := opts.SetAccessCredentialsFromEnv(rootCmd); err != nil {
		panic(err)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "s3-cleaner",
	Short:   "This tool finds the desired files in a bucket and cleans them",
	Long:    ``,
	Version: ver.GitVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		/*bannerBytes, _ := ioutil.ReadFile("banner.txt")
		banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))*/

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

		sess, err := aws.CreateSession(opts)
		if err != nil {
			logger.Error("an error occured while creating session", zap.Error(err))
			return err
		}

		// obtain S3 client with initialized session
		svc := s3.New(sess)

		logger.Debug("trying to find files on bucket", zap.String("bucketName", opts.BucketName),
			zap.String("region", opts.Region))

		if err := aws.GetAllFiles(svc, opts); err != nil {
			return err
		}

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
