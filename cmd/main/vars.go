package main

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

var (
	defaultLogger = logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	commonRequest = &requestDefinition{
		Subject: "user123",
		Object:  "api/v1/user",
		Action:  "get",
	}
)
