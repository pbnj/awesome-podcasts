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

	log "github.com/sirupsen/logrus"
	survey "gopkg.in/AlecAivazis/survey.v1"
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
	gen := flag.Bool("gen", false, "Generate README file")
	add := flag.Bool("add", false, "Add new podcast")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Please, specify one of the following flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}
	// 1. Read in JSON file
	b, err := ioutil.ReadFile("awesome-podcasts.json")
	if err != nil {
		log.Fatalf("could not read JSON file: %s", err)
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

	if *gen {
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
			log.Fatalf("could not create README file: %s", err)
		}
		defer f.Close()

		// 6. Create a buffered writer
		w := bufio.NewWriter(f)
		defer w.Flush()

		// 7. Write data to README
		err = t.ExecuteTemplate(w, "readme.md.tmpl", podcasts)
		if err != nil {
			log.Fatalf("could not write README file: %s", err)
		}
	}

	if *add {
		// TODO: implement add
		qs := []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "Podcast Name:"},
				Validate: survey.Required,
			},
			{
				Name:     "url",
				Prompt:   &survey.Input{Message: "Podcast URL:"},
				Validate: survey.Required,
			},
			{
				Name:     "desc",
				Prompt:   &survey.Input{Message: "Podcast Description:"},
				Validate: survey.Required,
			},
			{
				Name:     "category",
				Prompt:   &survey.Input{Message: "Podcast Category"},
				Validate: survey.Required,
			},
		}

		answers := struct {
			Name     string
			URL      string
			Desc     string
			Category string
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			log.Fatalf("could not prompt questions: %s", err)
		}

		for _, p := range podcasts {
			if p.Category == answers.Category {
				p.Pods = append(p.Pods, Pod{
					Name: answers.Name,
					URL:  answers.URL,
					Desc: answers.Desc,
				})
			}
			sort.Slice(p.Pods, func(i, j int) bool {
				return strings.ToUpper(p.Pods[i].Name) < strings.ToUpper(p.Pods[j].Name)
			})
		}

		b, err := json.MarshalIndent(podcasts, "", "  ")
		if err != nil {
			log.Fatalf("could not convert struct to JSON: %+v", err)
		}
		err = ioutil.WriteFile("awesome-podcasts.json", b, 0777)
		if err != nil {
			log.Fatalf("could not write JSON file: %+v", err)
		}
		log.WithFields(log.Fields{
			"Name":     answers.Name,
			"URL":      answers.URL,
			"Desc":     answers.Desc,
			"Category": answers.Category,
		}).Infoln("SUCCESS!")
	}
}
