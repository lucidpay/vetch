package main

type Routes []Route

var routes = Routes{
	Route{
		"Test",
		"GET",
		"/vetch/v1/test",
		Test,
	},
}
