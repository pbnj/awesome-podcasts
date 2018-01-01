.PHONY: readme

readme:
	go run main.go

docker:
	docker run --rm -it -v "$(shell pwd):/go" golang go run main.go
