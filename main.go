package main // import "github.com/petermbenjamin/awesome-podcasts"

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	jsonfile = "awesome-podcasts.json"
	yamlfile = "awesome-podcasts.yaml"
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
	b, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		logrus.Warnf("YAML file not found: %+s", err)
	}

	// Load data into Go structs
	var podcasts []Podcast
	err = yaml.Unmarshal(b, &podcasts)
	if err != nil {
		logrus.Errorf("could not unmarshal YAML: %+s", err)
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
		logrus.Warnf("could not marshal sorted JSON: %+v", err)
	}
	// Write JSON
	err = ioutil.WriteFile(jsonfile, marshaledBytes, 0644)
	if err != nil {
		logrus.Warnf("could not write sorted JSON: %+v", err)
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
		logrus.Fatalf("could not create README file: %s", err)
	}
	defer f.Close()

	// Create buffered writer
	w := bufio.NewWriter(f)
	defer w.Flush()

	// Write data out
	err = t.ExecuteTemplate(w, "readme.md.tmpl", podcasts)
	if err != nil {
		logrus.Fatalf("could not write README file: %s", err)
	}
}
