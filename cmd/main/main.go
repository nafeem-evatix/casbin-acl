package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"

	"github.com/nafeem-evatix/casbin-acl/internal/consts"
)

func main() {
	// getting the enforcer
	enforcer := getEnforcerMust()

	// loading all policies from db
	loadPolicy(enforcer)

	// before any policy is there
	checkAndPrintCommonRequest(enforcer)

	// adding policy
	addPolicy(enforcer, commonRequest)

	// checking after policy is added
	checkAndPrintCommonRequest(enforcer)

	// removing policy
	removePolicy(enforcer, commonRequest)

	// checking after policy is removed
	checkAndPrintCommonRequest(enforcer)

	// saving policy
	savePolicy(enforcer)
}

func checkAndPrintCommonRequest(enforcer *casbin.Enforcer) {
	allowed, err := checkIfRequestIsAllowed(enforcer, commonRequest)
	if err != nil {
		panic(err)
	}

	printPolicyAllowance(commonRequest, allowed)
}

func printPolicyAllowance(req *requestDefinition, isAllowed bool) {
	fmt.Printf(consts.AccessPrinterFormat,
		req.Subject,
		req.Object,
		req.Action,
		isAllowed)
}

func loadPolicy(enforcer *casbin.Enforcer) error {
	return enforcer.LoadPolicy()
}

func checkIfRequestIsAllowed(enforcer *casbin.Enforcer, req *requestDefinition) (bool, error) {
	return enforcer.Enforce(req.Subject, req.Object, req.Action)
}

func addPolicy(enforcer *casbin.Enforcer, req *requestDefinition) (bool, error) {
	return enforcer.AddPolicy(req.Subject, req.Object, req.Action)
}

func removePolicy(enforcer *casbin.Enforcer, req *requestDefinition) (bool, error) {
	return enforcer.RemovePolicy(req.Subject, req.Object, req.Action)
}

func savePolicy(enforcer *casbin.Enforcer) error {
	return enforcer.SavePolicy()
}
