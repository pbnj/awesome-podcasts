package main // import "github.com/petermbenjamin/awesome-podcasts"

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	jsonFile = "awesome-podcasts.json"
	yamlFile = "awesome-podcasts.yaml"
)

// Podcast represents the list of awesome podcasts
type Podcast struct {
	Category string `json:"category" yaml:"category"`
	Pods     []Pod  `json:"pods" yaml:"pods"`
	Subtitle string `json:"subtitle" yaml:"subtitle"`
}

// Pod represents a podcast object
type Pod struct {
	Desc string `json:"desc" yaml:"desc"`
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

func main() {

	// Read YAML file
	b, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		fmt.Println(fmt.Errorf("YAML file not found: %+s", err))
		os.Exit(1)
	}

	// Load data into Go structs
	var podcasts []Podcast
	err = yaml.Unmarshal(b, &podcasts)
	if err != nil {
		fmt.Println(fmt.Errorf("could not unmarshal YAML: %+s", err))
		os.Exit(1)
	}

	// Sort alphabetically by category
	sort.Slice(podcasts, func(i, j int) bool {
		return podcasts[i].Category < podcasts[j].Category
	})
	// Sort alphabetically by podcast, ignoring case-sensitivity
	for _, c := range podcasts {
		sort.Slice(c.Pods, func(i, j int) bool {
			return strings.ToUpper(c.Pods[i].Name) <
				strings.ToUpper(c.Pods[j].Name)
		})
	}

	// Generate JSON
	marshaledBytes, err := json.MarshalIndent(podcasts, "", "  ")
	if err != nil {
		fmt.Println(fmt.Errorf("could not marshal sorted JSON: %+v", err))
		os.Exit(1)
	}
	// Write JSON
	err = ioutil.WriteFile(jsonFile, marshaledBytes, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("could not write sorted JSON: %+v", err))
		os.Exit(1)
	}

	// helper functions
	funcMap := template.FuncMap{
		"dashed": func(word string) string {
			word = strings.ToLower(word)
			word = strings.Replace(word, " ", "-", -1)
			word = strings.Replace(word, "/", "", -1)
			return word
		},
		"titled": strings.Title,
	}

	// Generate README
	// Load README template
	paths := []string{
		filepath.Join("tmpl", "readme.md.tmpl"),
	}
	t := template.Must(template.
		New("main").
		Funcs(funcMap).
		ParseFiles(paths...))

	// Create file
	f, err := os.Create("README.md")
	if err != nil {
		fmt.Println(fmt.Errorf("could not create README file: %s", err))
	}
	defer f.Close()

	// Create buffered writer
	w := bufio.NewWriter(f)
	defer w.Flush()

	// Write data out
	err = t.ExecuteTemplate(w, "readme.md.tmpl", podcasts)
	if err != nil {
		fmt.Println(fmt.Errorf("could not write README file: %s", err))
		os.Exit(1)
	}
}
