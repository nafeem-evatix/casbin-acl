package main

type assignRoleDefinition struct {
	Subject, Role string
}

func newAssignRoleDefinition(subject, role string) *assignRoleDefinition {
	return &assignRoleDefinition{
		Subject: subject,
		Role:    role,
	}
}
