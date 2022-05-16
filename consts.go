package main

var ApisToRolesMap map[string]string

func mapAPIsToRoles() {
	ApisToRolesMap = make(map[string]string)
	ApisToRolesMap["Test"] = "ADMIN"
}
