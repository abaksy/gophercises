package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

func parseStory(filename string) (map[string]Chapter, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var json_data map[string]Chapter = make(map[string]Chapter)
	json.Unmarshal(data, &json_data)
	return json_data, nil
}

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

	for k, _ := range json_data {
		if k != "intro" {
			url := fmt.Sprintf("/%s", k)
			mux.Handle(url, storyHandler(json_data, k))
		}
	}

	mux.Handle("/", storyHandler(json_data, "intro"))

	fmt.Println("Listening on port 8080.....")
	http.ListenAndServe(":8080", mux)
}
