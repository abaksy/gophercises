package main

import (
	"encoding/json"
	"io/ioutil"
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
