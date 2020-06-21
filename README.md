# try-circle-ci

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![CircleCI](https://circleci.com/gh/kemokemo/try-circle-ci.svg?style=svg)](https://circleci.com/gh/kemokemo/try-circle-ci) [![codecov](https://codecov.io/gh/kemokemo/try-circle-ci/branch/master/graph/badge.svg)](https://codecov.io/gh/kemokemo/try-circle-ci) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/try-circle-ci)](https://goreportcard.com/report/github.com/kemokemo/try-circle-ci)

A repository to play around with using [CircleCI](https://circleci.com/).

## Install

You require the golang 1.13 or later. (because of the [testing.Init()](https://github.com/kemokemo/try-circle-ci/blob/492e2be35caa30474754c9f258e10615ccbe6185/main.go#L24))

```sh
go get kemokemo/try-circle-ci
```

## Usage

A simple greeter cli.

```sh
$ try-circle-ci kemokemo nekomimi
Hello kemokemo!
Hello nekomimi!
```

## About CI/CD

- CI: [CircleCI](https://circleci.com/gh/kemokemo/try-circle-ci)
- Coverage: [Codecov](https://codecov.io/gh/kemokemo/try-circle-ci)

## License

[MIT](https://github.com/kemokemo/try-circle-ci/blob/master/LICENSE)

## Authors

[kemokemo](https://github.com/kemokemo)
