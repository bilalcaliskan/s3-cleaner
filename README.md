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

## Required Steps
- Single command is mostly enough to prepare project, it will prompt you with some questions about your new project:
  ```shell
  $ make -s prepare-initial-project
  ```

## Additional nice-to-have steps
- If you want to build and publish Docker image:
  - Ensure `DOCKER_USERNAME` has been added as **repository secret on GitHub**
  - Ensure `DOCKER_PASSWORD` has been added as **repository secret on GitHub**
  - Uncomment **line 178** to **line 185** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 32** to **line 50** in [build/package/.goreleaser.yaml](build/package/.goreleaser.yaml)
- If you want to enable https://sonarcloud.io/ integration:
  - Ensure your created repository from that template has been added to https://sonarcloud.io/
  - Ensure `SONAR_TOKEN` has been added as **repository secret** on GitHub
  - Ensure `SONAR_TOKEN` has been added as **dependabot secret** on GitHub
  - Uncomment **line 149** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 94** to **line 123** in [.github/workflows/push.yml](.github/workflows/push.yml)
- If you want to create banner:
  - Generate a banner from [here](https://devops.datenkollektiv.de/banner.txt/index.html) and place it inside of [build/ci](build/ci) directory
  - Uncomment **line 30** and **line 31** in [cmd/root.go](cmd/root.go)
  - Run `go get -u github.com/dimiro1/banner`
- If you want to release as Homebrew Formula:
  - At first, you must have a **formula repository** like https://github.com/bilalcaliskan/homebrew-tap
  - Ensure `TAP_GITHUB_TOKEN` has been added as **repository secret** on GitHub
  - Uncomment **line 198** in [.github/workflows/push.yml](.github/workflows/push.yml)
  - Uncomment **line 70** to **line 80** in [build/package/.goreleaser.yaml](build/package/.goreleaser.yaml)

## Used Libraries
- [spf13/cobra](https://github.com/spf13/cobra)
- [stretchr/testify](https://github.com/stretchr/testify)
- [go.uber.org/zap](https://go.uber.org/zap)

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
