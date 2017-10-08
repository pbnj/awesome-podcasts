package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type AP struct {
	AwesomePodcasts []PodcastGroup `json:"podcasts"`
}

// PodcastGroup represents a single category of awesome podcast
type PodcastGroup struct {
	Category string    `json:"category"`
	Link     string    `json:"link"`
	SubTitle string    `json:"subtitle"`
	Pods     []Podcast `json:"pods"`
}

// Podcast represents a single podcast object
type Podcast struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Desc string `json:"desc"`
}

const readmeTemplate = `# Awesome Podcasts

> 😎 Curated list of awesome programming podcasts  [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

---

<details>

<summary>Table of Contents</summary>

{{range .AwesomePodcasts}}- [{{.Category}}](#{{.Link}})
{{end}}
</details>

---

{{range .AwesomePodcasts}}## {{.Category}}

> {{.SubTitle}}

{{range .Pods}}- [{{.Name}}]({{.URL}}) - {{.Desc}}
{{end}}
{{end}}
`

func main() {
	// First, read in the JSON file
	jsBytes, readErr := ioutil.ReadFile("awesome-podcasts.json")
	if readErr != nil {
		log.Fatalf("could not read JSON file: %+v\n", readErr)
	}

	// Second, unmarshal data from JSON into Go data structures
	var awesomePodcasts AP
	unmarshErr := json.Unmarshal(jsBytes, &awesomePodcasts)
	if unmarshErr != nil {
		log.Fatalf("could not unmarshal data: %+v\n", unmarshErr)
	}

	// log.Printf("Data: %+v\n", awesomePodcasts)
	// Third, create README file bufer
	f, createErr := os.Create("README.md")
	if createErr != nil {
		log.Fatalf("could not create README file: %+v\n", createErr)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	t := template.Must(template.New("tmpl").Parse(readmeTemplate))
	templateErr := t.Execute(w, awesomePodcasts)
	if templateErr != nil {
		log.Fatalf("could not merge data into template: %+v\n", templateErr)
	}
	w.Flush()
}
