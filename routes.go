package main

type Routes []Route

var routes = Routes{
	Route{
		"Health",
		"POST",
		"/vetch/v1/health",
		Health,
	},
}
