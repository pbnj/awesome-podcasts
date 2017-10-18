.PHONY: readme

readme:
	go run main.go

html: readme
	open docs/index.html
