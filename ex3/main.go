package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

func storyHandler(json_data map[string]Chapter, arc string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/story.html"))
		tmpl.Execute(w, json_data[arc])
	}
}

func main() {
	fileNameFlag := flag.String("f", "gopher.json", "Filename to read story from")
	var filename string = *fileNameFlag

	// http.Handle("/story", storyHandler(filename))
	mux := http.NewServeMux()

	json_data, err := parseStory(filename)
	if err != nil {
		panic(err)
	}

	for k := range json_data {
		if k != "intro" {
			url := fmt.Sprintf("/%s", k)
			mux.Handle(url, storyHandler(json_data, k))
		}
	}

	mux.Handle("/", storyHandler(json_data, "intro"))

	fmt.Println("Listening on port 8080.....")
	http.ListenAndServe(":8080", mux)
}
