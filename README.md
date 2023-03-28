# s3-cleaner
[![CI](https://github.com/bilalcaliskan/s3-cleaner/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/s3-cleaner/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/s3-cleaner)](https://hub.docker.com/r/bilalcaliskan/s3-cleaner/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/s3-cleaner)](https://goreportcard.com/report/github.com/bilalcaliskan/s3-cleaner)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-cleaner&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-cleaner)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-cleaner&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-cleaner)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-cleaner&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-cleaner)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-cleaner&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-cleaner)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_s3-cleaner&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_s3-cleaner)
[![Release](https://img.shields.io/github/release/bilalcaliskan/s3-cleaner.svg)](https://github.com/bilalcaliskan/s3-cleaner/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/s3-cleaner)](https://github.com/bilalcaliskan/s3-cleaner)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

It is a tool for finding desired files in a S3 bucket and cleans them with a rule set.

You can provide access credentials of your AWS account with below environment variables or CLI flags. Keep in mind that command line flags
will override environment variables if you set both of them:
```
"--accessKey" CLI flag or "AWS_ACCESS_KEY" environment variable
"--secretKey" CLI flag or "AWS_SECRET_KEY" environment variable
"--region" CLI flag or "AWS_REGION" environment variable
"--bucketName" CLI flag or "AWS_BUCKET_NAME" environment variable
```

## Configuration
```
Usage:
  s3-cleaner start [flags]

Flags:
      --autoApprove             Skip interactive approval (default false)
      --dryRun                  specifies that if you just want to see what to delete or completely delete them all (default false)
      --fileExtensions string   selects the files with defined extensions to clean from target bucket, "" means all files (default "")
  -h, --help                    help for start
      --keepLastNFiles int      defines how many of the files to skip deletion in specified criteria, 0 means clean them all (default 1)
      --maxFileSizeInMb int     maximum size in mb to clean from target bucket, 0 means no upper limit (default 15)
      --minFileSizeInMb int     minimum size in mb to clean from target bucket, 0 means no lower limit (default 10)
      --sortBy string           defines the ascending order in the specified criteria, valid options are "lastModificationDate" and "size" (default "lastModificationDate")

Global Flags:
      --accessKey string        access key credential to access S3 bucket, this value also can be passed via "AWS_ACCESS_KEY" environment variable (default "")
      --secretKey string        secret key credential to access S3 bucket, this value also can be passed via "AWS_SECRET_KEY" environment variable (default "")
      --region string           region of the target bucket on S3, this value also can be passed via "AWS_REGION" environment variable (default "")
      --bucketName string       name of the target bucket on S3, this value also can be passed via "AWS_BUCKET_NAME" environment variable (default "")
      --fileNamePrefix string   folder name of target bucket objects, means it can be used for folder-based object grouping buckets (default "")
  -v, --verbose                 verbose output of the logging library (default false)
```

> **WARNING:**
> Please note that environmental flags for accessing AWS (--accessKey, --secretKey etc.) takes precedence over environment variables for these flags (AWS_ACCESS_KEY, AWS_SECRET_KEY etc.)

## Installation

### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/s3-cleaner/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./s3-cleaner start --accessKey=xxxxx --secretKey=xxxxx --region=xxxxx --bucketName=xxxxx
```

or by providing environment variables for access credentials:
```shell
$ AWS_ACCESS_KEY=xxxxx AWS_SECRET_KEY=xxxxx AWS_REGION=xxxxx AWS_BUCKET_NAME=xxxxx ./s3-cleaner start
```

### Homebrew
This project can be installed with [Homebrew](https://brew.sh/):
```shell
$ brew tap bilalcaliskan/tap
$ brew install bilalcaliskan/tap/s3-cleaner
```

Then similar to binary method, you can run it by calling below command:
```shell
$ s3-cleaner start --accessKey=xxxxx --secretKey=xxxxx --region=xxxxx --bucketName=xxxxx
```

or by providing environment variables for access credentials:
```shell
$ AWS_ACCESS_KEY=xxxxx AWS_SECRET_KEY=xxxxx AWS_REGION=xxxxx AWS_BUCKET_NAME=xxxxx ./s3-cleaner start
```

## Development
This project requires below tools while developing:
- [Golang 1.20](https://golang.org/doc/go1.20)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

Simply run below command to prepare your development environment:
```shell
$ python3 -m venv venv
$ source venv/bin/activate
$ pip3 install pre-commit
$ pre-commit install -c build/ci/.pre-commit-config.yaml
```
