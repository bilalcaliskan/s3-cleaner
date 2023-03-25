package cleaner

import (
	"errors"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	start "github.com/bilalcaliskan/s3-cleaner/cmd/start/options"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog"
)

func getProperObjects(startOpts *start.StartOptions, allFiles *s3.ListObjectsOutput, logger zerolog.Logger) (res []*s3.Object) {
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

	return res
}

func sortObjects(slice []*s3.Object, startOpts *start.StartOptions) {
	switch startOpts.SortBy {
	case "lastModificationDate":
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].LastModified.Before(*slice[j].LastModified)
		})
	case "size":
		sort.Slice(slice, func(i, j int) bool {
			return *slice[i].Size < *slice[j].Size
		})
	}
}

func checkLength(targetObjects []*s3.Object) error {
	if len(targetObjects) == 0 {
		return errors.New("no deletable file found on the target bucket")
	}

	return nil
}

func promptDeletion(startOpts *start.StartOptions, logger zerolog.Logger, keys []string) error {
	if !startOpts.AutoApprove {
		logger.Info().Any("files", keys).Msg("these files will be removed if you approve:")

		prompt := promptui.Prompt{
			Label:     "Delete Files? (y/N)",
			IsConfirm: true,
			Validate: func(s string) error {
				if len(s) == 1 {
					return nil
				}

				return errors.New("invalid input")
			},
		}

		if res, err := prompt.Run(); err != nil {
			if strings.ToLower(res) == "n" {
				return errors.New("user terminated the process")
			}

			return errors.New("invalid input")
		}
	}

	return nil
}
