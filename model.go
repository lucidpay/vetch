package main

import (
	"net/http"
)

type Config struct {
	LocaltimeToUTC			int     `yaml:"localtime_to_utc"`		// 2- Micro service local time diff
	MSPort					int		`yaml:"msport"`					// 3- Micro service listening port
	JWTKey				 string		`yaml:"jwt_key"`				// 4- JWT secret key
	IsSecure			   bool		`yaml:"secure_via_jwt"`			// 5- Do APIs require JWT session?
}

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

