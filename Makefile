.PHONY: readme

readme:
	go run main.go

html: readme
	open docs/index.html

docker:
	docker run --rm -it -v "$(shell pwd):/go" golang go run main.go
