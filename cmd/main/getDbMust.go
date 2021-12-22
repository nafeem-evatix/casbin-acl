package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/nafeem-evatix/casbin-acl/internal/consts"
)

func getDbMust() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(consts.DbPath),
		&gorm.Config{
			Logger: defaultLogger,
		},
	)

	if err != nil {
		panic(err)
	}

	return db
}
