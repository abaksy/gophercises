package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"urlshort"
)

func help() string {
	helpstring := `
USAGE: ./urlshort -f <redirect URLs file> [-h]
	
File can be a valid YAML or valid JSON file.

YAML File format: 
- path: /wiki
  url: https://wikipedia.com
- path: /urlshort
  url: https://github.com/gophercises/urlshort/

JSON File format:
[
	{
		"path": "/wiki"
  		"url" : "https://wikipedia.com"
	},
	{
		"path": "/urlshort"
  		"url" : "https://github.com/gophercises/urlshort/"
	}
]
`
	return helpstring
}

func buildHandler(fileName string) (http.HandlerFunc, error) {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var fileParts []string = strings.Split(fileName, ".")
	var fileExtension string = fileParts[len(fileParts)-1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if fileExtension == "json" {
		handler, err := urlshort.JSONHandler([]byte(data), mapHandler)
		if err != nil {
			return nil, err
		}
		return handler, nil
	} else if fileExtension == "yml" {
		handler, err := urlshort.YAMLHandler([]byte(data), mapHandler)
		if err != nil {
			return nil, err
		}
		return handler, nil
	}
	return nil, errors.New("invalid file format")
}

func main() {

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var fileFlag *string = flag.String("f", "", "File containing redirect URLs")
	helpFlag := flag.Bool("h", false, "Print command help")
	flag.Parse()

	if *helpFlag || (!(*helpFlag) && *fileFlag == "") {
		helpString := help()
		fmt.Println(helpString)
		os.Exit(0)
	}
	var fileName string = *fileFlag
	handler, err := buildHandler(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
