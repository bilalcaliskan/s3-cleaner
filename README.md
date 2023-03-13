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

## Configuration
```
This tool finds the desired files in a bucket and cleans them

Usage:
  s3-cleaner [flags]

Flags:
      --accessKey string    access key credential to access S3 bucket
      --bucketName string   name of the target bucket on S3
  -h, --help                help for s3-cleaner
      --region string       region of the target bucket on S3
      --secretKey string    secret key credential to access S3 bucket
  -v, --version             version for s3-cleaner

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
