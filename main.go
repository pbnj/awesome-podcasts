package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// AwesomePodcasts reperesents a slice of Podcast Categories
type AwesomePodcasts struct {
	Categories []Category `json:"categories"`
}

// Category represents a single category of awesome podcast
type Category struct {
	Name     string    `json:"name"`
	SubTitle string    `json:"subtitle"`
	Pods     []Podcast `json:"pods"`
}

// Podcast represents a single podcast object
type Podcast struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Desc string `json:"desc"`
}

func main() {
	// First, read in the JSON file
	jsBytes, readErr := ioutil.ReadFile("awesome-podcasts.json")
	if readErr != nil {
		log.Fatalf("could not read JSON file: %+v\n", readErr)
	}

	// Second, unmarshal data from JSON into Go data structures
	var awesomePodcasts AwesomePodcasts
	unmarshErr := json.Unmarshal(jsBytes, &awesomePodcasts)
	if unmarshErr != nil {
		log.Fatalf("could not unmarshal data: %+v\n", unmarshErr)
	}

	// Third, sort categories alphabetically
	sort.Slice(awesomePodcasts.Categories, func(i, j int) bool {
		return awesomePodcasts.Categories[i].Name < awesomePodcasts.Categories[j].Name
	})
	// Fourth, sort podcasts in each category alphabetically
	for _, category := range awesomePodcasts.Categories {
		sort.Slice(category.Pods, func(i, j int) bool {
			return strings.ToUpper(category.Pods[i].Name) < strings.ToUpper(category.Pods[j].Name)
		})
	}

	// Fifth, setup templates & custom template functions
	paths := []string{
		filepath.Join("templates", "README.md"),
	}

	funcMap := template.FuncMap{
		// used for anchor/href tags
		"dashed": func(word string) string {
			word = strings.ToLower(word)
			word = strings.Replace(word, " ", "-", -1)
			word = strings.Replace(word, "/", "", -1)
			return word
		},
		"titled": strings.Title,
	}

	// last, generate README
	rf, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("could not create README file: %+v\n", err)
	}
	defer rf.Close()
	rw := bufio.NewWriter(rf)
	defer rw.Flush()

	err = t.ExecuteTemplate(rw, "README.md", awesomePodcasts)
	if err != nil {
		log.Fatalf("could not create generate README: %+v\n", err)
	}
}
