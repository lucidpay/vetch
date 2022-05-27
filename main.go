package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var requestLogger = logrus.New().WithFields(logrus.Fields{
	"appName": "vetch",
})

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
		requestLogger.Fatal(err)
	}

	err = yaml.Unmarshal(source, &Config_)
	if err != nil {
		requestLogger.Fatal(err)
	}

	requestLogger.Infoln("==>> Config unmarshaled properly")
}

func main() {
	// Setup correctly the GOPATH in the environement
	if goPath := os.Getenv("GOPATH"); goPath == "" {
		os.Setenv("GOPATH", defaultGOPATH())
	}

	router := NewRouter()

	requestLogger.Infoln("==>> Micro service started. Listening on port ", strconv.Itoa(Config_.MSPort))
	requestLogger.Fatal(http.ListenAndServe(":"+strconv.Itoa(Config_.MSPort), router))
}
