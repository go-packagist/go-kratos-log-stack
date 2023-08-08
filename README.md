# go-kratos-log-stack

[![Go Version](https://badgen.net/github/release/go-packagist/go-kratos-log-stack/stable)](https://github.com/go-packagist/go-kratos-log-stack/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/go-kratos-log-stack)](https://pkg.go.dev/github.com/go-packagist/go-kratos-log-stack)
[![codecov](https://codecov.io/gh/go-packagist/go-kratos-log-stack/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/go-kratos-log-stack)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/go-kratos-log-stack)](https://goreportcard.com/report/github.com/go-packagist/go-kratos-log-stack)
[![tests](https://github.com/go-packagist/go-kratos-log-stack/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/go-kratos-log-stack/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/go-kratos-log-stack
```

## Usage

```go
package main

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	slack "github.com/go-packagist/go-kratos-log-stack"
)

func main() {
	logger := slack.New([]log.Logger{
		log.NewStdLogger(os.Stdout),
		log.NewStdLogger(os.Stdout),
	})

	log.SetLogger(logger)

	log.Info("test", "test")
	// Print:
	// INFO msg=testtest
	// INFO msg=testtest
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.