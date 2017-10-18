package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
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

const htmlTemplate = `
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Awesome Podcasts</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
	<link href="https://fonts.googleapis.com/css?family=Pacifico" rel="stylesheet">
</head>
<style>

body {
	padding-top: 4.5rem;
}

.podcast-category {
	font-family: monospace;
}

.podcast-subtitle {
	font-family: monospace;
}

a:hover {
	text-decoration: none;
}

a.navbar-brand {
	font-family: 'Pacifico', cursive;
}

.hvr-underline-from-center {
	display: inline-block;
	vertical-align: middle;
	-webkit-transform: perspective(1px) translateZ(0);
	transform: perspective(1px) translateZ(0);
	box-shadow: 0 0 1px transparent;
	position: relative;
	overflow: hidden;
}
.hvr-underline-from-center:before {
	content: "";
	position: absolute;
	z-index: -1;
	left: 50%;
	right: 50%;
	bottom: 0;
	/*background: tomato;*/
	animation: colorchange 2s infinite;
	-webkit-animation: colorchange 2s infinite;
	-moz-animation: colorchange 2s infinite;
	height: 4px;
	-webkit-transition-property: left, right;
	transition-property: left, right;
	-webkit-transition-duration: 0.3s;
	transition-duration: 0.3s;
	-webkit-transition-timing-function: ease-out;
	transition-timing-function: ease-out;
}
.hvr-underline-from-center:hover:before, .hvr-underline-from-center:focus:before, .hvr-underline-from-center:active:before {
	left: 0;
	right: 0;
}

@keyframes colorchange
{
  0%   {background: red;}
  25%  {background: yellow;}
  50%  {background: blue;}
  75%  {background: green;}
  100% {background: red;}
}

@-webkit-keyframes colorchange /* Safari and Chrome - necessary duplicate */
{
  0%   {background: red;}
  25%  {background: yellow;}
  50%  {background: blue;}
  75%  {background: green;}
  100% {background: red;}
}
</style>
<body>

	<nav class="navbar fixed-top navbar-expand-lg navbar-light bg-light">
		<a class="navbar-brand" href="#">Awesome Podcasts</a>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarNavDropdown">
			<ul class="navbar-nav">
				<li class="nav-item">
				<a class="nav-link" href="#">Top 👆</a>
				</li>
				<li class="nav-item">
				<a class="nav-link" href="https://github.com/petermbenjamin/awesome-podcasts">GitHub 🐙</a>
				</li>
				<li class="nav-item dropdown">
				<li class="nav-item dropdown">
				<a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
					Categories
				</a>
				<div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
					{{range .AwesomePodcasts}}<a class="dropdown-item" href="#{{.Link}}">{{.Category}}</a>
					{{end}}
				</div>
				</li>
			</ul>
		</div>
	</nav>

	{{range .AwesomePodcasts}}
	<div class="container">
		<a href="#{{.Link}}" name="{{.Link}}" class="hvr-underline-from-center">
			<h4 class="podcast-category"># {{.Category}}</h4>
		</a>
		<blockquote class="podcast-subtitle">> {{.SubTitle}}</blockquote>
	</div>
	<div class="container">
		<table class="table table-bordered">
			<thead>
				<tr>
					<th>Name</th>
					<th>Description</th>
				</tr>
			</thead>
			<tbody>
			{{range .Pods}}<tr>
				<td><a href="{{.URL}}">{{.Name}}</a></td>
				<td>{{.Desc}}</td>
			</tr>
			{{end}}
			</tbody>
		</table>
	</div>
	{{end}}
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</body>
</html>
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

	// Third, sort categories alphabetically
	sort.Slice(awesomePodcasts.AwesomePodcasts, func(i, j int) bool {
		return awesomePodcasts.AwesomePodcasts[i].Category < awesomePodcasts.AwesomePodcasts[j].Category
	})

	// Fourth, sort podcasts in each category alphabetically
	for _, category := range awesomePodcasts.AwesomePodcasts {
		sort.Slice(category.Pods, func(i, j int) bool {
			return strings.ToUpper(category.Pods[i].Name) < strings.ToUpper(category.Pods[j].Name)
		})
	}

	// Last, create README file bufer, write content, flush buffer and close file handler
	rf, createErr := os.Create("README.md")
	if createErr != nil {
		log.Fatalf("could not create README file: %+v\n", createErr)
	}
	defer rf.Close()

	rw := bufio.NewWriter(rf)
	rt := template.Must(template.New("tmpl").Parse(readmeTemplate))
	templateErr := rt.Execute(rw, awesomePodcasts)
	if templateErr != nil {
		log.Fatalf("could not merge data into template: %+v\n", templateErr)
	}
	rw.Flush()

	hf, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatalf("could not create index.html file: %+v\n", err)
	}
	defer hf.Close()

	hw := bufio.NewWriter(hf)
	ht := template.Must(template.New("tmpl").Parse(htmlTemplate))
	err = ht.Execute(hw, awesomePodcasts)
	if err != nil {
		log.Fatalf("could not merge data into html template: %+v\n", err)
	}
	hw.Flush()
}
