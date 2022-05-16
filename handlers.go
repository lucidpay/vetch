package main

import (
	"bytes"
	crypto "crypto/ed25519"
	"fmt"
	"io"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	// Sign a message with ed25519:
	// check crypto/ed25519 documentation here https://pkg.go.dev/crypto/ed25519
	priv_key := "edsk4CLwrvzBrd2njFMr6EgGcrYrpptZu8h8hbYEyFVd2kQxHVC6t3hhhhhhhhhh"
	signature := crypto.Sign([]byte(priv_key), []byte("message"))
	fmt.Println("signature: ", signature)

	// Sending an http request
	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://node-url", bytes.NewBuffer(nil))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("The HTTP request to the node failed with error:", err)
		http.Error(w, "{\"status\":\"The HTTP request to the node failed\"}", http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("{\"status\":\"SUCCESS\"}"))
}
