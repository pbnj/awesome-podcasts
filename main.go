package main // import "github.com/petermbenjamin/awesome-podcasts"

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	AWESOMEPODCASTJSONFILE = "awesome-podcasts.json"
)

type Podcast struct {
	Category string `json:"category"`
	Pods     []Pod  `json:"pods"`
	Subtitle string `json:"subtitle"`
}

type Pod struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	generate := flag.Bool("gen", false, "Generate README file")
	format := flag.Bool("fmt", false, "Format JSON file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Please, specify one of the following flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	// 1. Read in JSON file
	b, err := ioutil.ReadFile(AWESOMEPODCASTJSONFILE)
	if err != nil {
		logrus.Warnf("JSON file not found [%s]: %s", AWESOMEPODCASTJSONFILE, err)
	}

	// 2. Load in data into Go struct
	var podcasts []Podcast
	json.Unmarshal(b, &podcasts)

	// 3a. sort alphabetically by category
	sort.Slice(podcasts, func(i, j int) bool {
		return podcasts[i].Category < podcasts[j].Category
	})
	// 3b. sort alphabetically by podcast
	for _, c := range podcasts {
		sort.Slice(c.Pods, func(i, j int) bool {
			return strings.ToUpper(c.Pods[i].Name) < strings.ToUpper(c.Pods[j].Name)
		})
	}

	if *format {
		marshaledBytes, err := json.MarshalIndent(podcasts, "", "  ")
		if err != nil {
			logrus.Warnf("could not marshal podcasts into sorted json: %+v", err)
		}

		err = ioutil.WriteFile(AWESOMEPODCASTJSONFILE, marshaledBytes, 0644)
		if err != nil {
			logrus.Warnf("could not write sorted json file: %+v", err)
		}
	}

	if *generate {
		// 4a. Set up template path
		paths := []string{
			filepath.Join("tmpl", "readme.md.tmpl"),
		}
		// 4b. Set up helper functions
		funcMap := template.FuncMap{
			"dashed": func(word string) string {
				word = strings.ToLower(word)
				word = strings.Replace(word, " ", "-", -1)
				word = strings.Replace(word, "/", "", -1)
				return word
			},
			"titled": strings.Title,
		}

		// 4c. Load in template
		t := template.Must(template.New("main").Funcs(funcMap).ParseFiles(paths...))

		// 5. Create file
		f, err := os.Create("README.md")
		if err != nil {
			logrus.Fatalf("could not create README file: %s", err)
		}
		defer f.Close()

		// 6. Create a buffered writer
		w := bufio.NewWriter(f)
		defer w.Flush()

		// 7. Write data to README
		err = t.ExecuteTemplate(w, "readme.md.tmpl", podcasts)
		if err != nil {
			logrus.Fatalf("could not write README file: %s", err)
		}
	}
}
