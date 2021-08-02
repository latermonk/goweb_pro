package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	fmt.Println("==============Begining====================")



	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "https://www.baidu.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)


	fmt.Println("===========Ending =================")

}
