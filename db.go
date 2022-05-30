package main

import (
	"fmt"

	"github.com/fjl/go-couchdb"
	couch "github.com/fjl/go-couchdb"
)

func NewDbObject() *couchdb.DB {
	client, err := couch.NewClient(
		fmt.Sprintf("http://%s:%s@%s:%s",
			Config_.DatabaseConstants.UserName,
			Config_.DatabaseConstants.Password,
			Config_.DatabaseConstants.Server,
			Config_.DatabaseConstants.Port), nil)
	if err != nil {
		requestLogger.Fatal(err)
	}

	// create if db doesn't not exist, returns db in both cases
	db, err := client.CreateDB(Config_.DatabaseConstants.Name)
	if err != nil {
		requestLogger.Infoln("didn't create db: ", err)
	}

	return db
}
