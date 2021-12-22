package main

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func getAdapterMust() *gormadapter.Adapter {
	db := getDbMust()
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	return adapter
}
