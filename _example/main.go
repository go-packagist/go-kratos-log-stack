package main

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	stack "github.com/go-packagist/go-kratos-log-stack"
)

func main() {
	logger := stack.New([]log.Logger{
		log.NewStdLogger(os.Stdout),
		log.NewStdLogger(os.Stdout),
	})

	log.SetLogger(logger)

	log.Info("test", "test")
}
