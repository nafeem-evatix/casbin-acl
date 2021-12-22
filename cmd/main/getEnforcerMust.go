package main

import (
	"github.com/casbin/casbin/v2"

	"github.com/nafeem-evatix/casbin-acl/internal/consts"
)

func getEnforcerMust() *casbin.Enforcer {
	adapter := getAdapterMust()
	enforcer, err := casbin.NewEnforcer(consts.ConfigurationFilePath, adapter)
	if err != nil {
		panic(err)
	}

	return enforcer
}
