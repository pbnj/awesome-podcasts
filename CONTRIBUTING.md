# CONTRIBUTING

Thank you for your interest in contributing and improving this project.

## Prerequisites

- [Go](https://golang.org)
- Or [Docker](https://www.docker.com)
- Or [c9.io](https://c9.io) account

## Steps

- Add the podcast information (e.g. `name`, `url`, and `desc`) under the correct category in the [`awesome-podcasts.json`](awesome-podcasts.json) JSON file.
- If you have go installed locally: `make readme`
	- If you are on a Mac & want to preview the HTML: `make html`
- If you have docker installed locally: `make docker`
- If you don't have go or docker, then spin up a generic workspace on [c9.io](https://c9.io) (it comes with go out of the box) & run: `go run main.go`

