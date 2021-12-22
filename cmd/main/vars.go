package main

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/nafeem-evatix/casbin-acl/internal/consts"
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
		Subject: consts.User123,
		Object:  consts.Data1,
		Action:  consts.Action,
	}

	adminRequest = &requestDefinition{
		Subject: consts.AdminRole,
		Object:  consts.Data1,
		Action:  consts.Action,
	}
)
