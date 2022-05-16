package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"gopkg.in/yaml.v2"
)

func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}

	if home := os.Getenv(env); home != "" {
		def := filepath.Join(home, "go")
		if filepath.Clean(def) == filepath.Clean(runtime.GOROOT()) {
			// Don't set the default GOPATH to GOROOT
			// as that will trigger warnings from the go tool
			return "" // ==>> EXITPOINT
		}
		return def // ==>> EXITPOINT
	}
	return "" // ==>> EXITPOINT
}

var Config_ Config

func init() {
	filename := "./vetch.yaml"
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &Config_)
	if err != nil {
		panic(err)
	}

	log.Printf("==>> Config unmarshaled properly")
}

func main() {
	// Setup correctly the GOPATH in the environement
	if goPath := os.Getenv("GOPATH"); goPath == "" {
		os.Setenv("GOPATH", defaultGOPATH())
	}

	router := NewRouter()

	log.Printf("==>> Micro service started. Listening on port %v", strconv.Itoa(Config_.MSPort))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config_.MSPort), router))
}
