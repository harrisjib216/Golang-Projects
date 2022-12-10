package main

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse the path
		path := r.URL.Path

		// if we find a match
		if dest, ok := pathsToUrls[path]; ok {
			fmt.Println("Navigating from:", path)
			fmt.Println("Navigating to:", dest)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		
		// else
		fmt.Println("Path not stored :(", "\nNavigating to the home page")
		fallback.ServeHTTP(w, r)
	}
}

type pathURL struct {
	path string `yaml:"path"`
	url string `yaml:"url"`
}

func YAMLHandler(yamlBtyes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse the yaml
	var pathURLS []pathURL

	err := yaml.Unmarshal(yamlBtyes, &pathURLS)
	if err != nil {
		return nil, err
	}

	// convert yaml array into map
	pathsToUrls := make(map[string]string{})
	for _, pu := range pathURLS {
		pathsToUrls[pu.path] = pu.url
	}

	// return a map hanlder
	return MapHandler(pathsToUrls, fallback), nil
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":80", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", errorPage)
	return mux
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}