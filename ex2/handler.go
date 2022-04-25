package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type URL struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func parseYAML(data []byte) ([]URL, error) {
	var yamlObject []URL
	err := yaml.Unmarshal(data, &yamlObject)
	if err != nil {
		return nil, err
	}
	return yamlObject, nil
}

func parseJSON(data []byte) ([]URL, error) {
	var jsonObject []URL
	err := json.Unmarshal(data, &jsonObject)
	if err != nil {
		return nil, err
	}
	return jsonObject, nil
}

func buildMap(urls []URL) map[string]string {
	var URLMap map[string]string = make(map[string]string)
	for _, v := range urls {
		URLMap[v.Path] = v.Url
	}
	return URLMap
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYML, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	URLMap := buildMap(parsedYML)
	return MapHandler(URLMap, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	URLMap := buildMap(parsedJSON)
	return MapHandler(URLMap, fallback), nil
}
