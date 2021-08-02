package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	fmt.Println("==============Begining====================")

	url := "https://www.baidu.com"

	/* Sugar log 的使用方式*/
/*
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
*/


	/* 生产环境 zap log 的使用方式*/
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	//logger.Info("failed to fetch URL",
	//	// Structured context as strongly typed Field values.
	//	zap.String("url", url),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second),
	//)

	logger.Error( "panic",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
		)
	// 如何打印颜色出来？



	fmt.Println("===========Ending =================")

}
