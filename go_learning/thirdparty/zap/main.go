package main

import (
	"time"
	"go.uber.org/zap"
)


func main() {

	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"./zap.log"}
	zapConfig.ErrorOutputPaths = []string{"./zap.log"}
	logger, _ := zapConfig.Build()
	// defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "https://dogkeeper.daemoncoder.com/"),
		zap.Duration("backoff", time.Millisecond * 1500),
	)

	time.Sleep(time.Second * 10)

	logger.Warn("failed to fetch URL",
		zap.String("url", "https://dogkeeper.daemoncoder.com/"),
		zap.Duration("backoff", time.Millisecond * 1500),
	)

	logger, _ = zapConfig.Build()

	logger.Error("failed to fetch URL",
		zap.String("url", "https://dogkeeper.daemoncoder.com/"),
		zap.Duration("backoff", time.Millisecond * 1500),
	)
}