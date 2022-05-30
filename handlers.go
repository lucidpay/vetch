package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func Health(w http.ResponseWriter, r *http.Request) {
	var doc Doc
	doc.Name = "testing"
	key := uuid.New().String()

	// put to database
	db := NewDbObject()
	_, err := db.Put(key, doc, "")
	if err != nil {
		requestLogger.Errorln("db failure: ", err)
		http.Error(w, "{\"status\":\"db failure\"}", http.StatusInternalServerError)
		return
	}

	// get from database
	var doc2 Doc
	err = db.Get(key, &doc2, nil)
	if err != nil {
		requestLogger.Errorln("db failure: ", err)
		http.Error(w, "{\"status\":\"db failure\"}", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("{\"key\":\"%s\",\"name\":\"%s\",\"rev\":\"%s\"}",
		key, doc2.Name, doc2.Rev))
}
