package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// zap package requires creating logger object first, then using log function.
	// NewProduction builds a sensible production Logger that writes InfoLevel and
	// above logs to standard error as JSON.
	// logger, _ := zap.NewProduction()
	// The reason why we use zap.NewProductionConfig is that we want to set the timestamp format.
	//
	// zap also provides `zap.NewExample()`, `zap.NewDevelopment()` to quickly create a logger, and the logger
	// created by different methods has different settings. Example is suitable for use in test code, Development is
	// used in the development environment, and Production is used in the production environment.
	// If you want to customize the logger, you can call the `zap.New()` method to create it.
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		// flushes buffer, if any
		_ = logger.Sync()
	}(logger)

	// the output is like this:
	// {"level":"info","ts":"2023-10-24T11:06:18+08:00","caller":"uber-zap-demo/demo-1.go:30","msg":"failed to fetch URL","url":"http://marmotedu.com","attempt":3,"backoff":1}
	url := "http://marmotedu.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// why use sugar logger? It supports Info, Infow, Infof forms.
	// The meaning of w in infow ? With Compared with info, infow does not need to use zap.String, zap.Int, zap.Duration and other functions.
	// When we have higher performance requirements for logs, we can use Logger instead of SugaredLogger. Logger
	// has better performance and fewer memory allocations. Logger does not use interface and reflection, and Logger only
	// supports structured logs, so when using Logger, you need to specify the specific type and key-value format log field.
	// You can call logger.Sugar() to create SugaredLogger. The use of SugaredLogger is simpler than Logger, but the performance
	// is about 50% lower than Logger
	//
	// the output is like this:
	// {"level":"info","ts":"2023-10-24T11:06:18+08:00","caller":"uber-zap-demo/demo-1.go:41","msg":"failed to fetch URL","url":"http://marmotedu.com","attempt":3,"backoff":1}
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	// {"level":"info","ts":"2023-10-24T11:06:18+08:00","caller":"uber-zap-demo/demo-1.go:45","msg":"failed to fetch URL: http://marmotedu.com"}
	sugar.Infof("failed to fetch URL: %s", url)
}
