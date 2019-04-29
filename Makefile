.DEFAULT_GOAL := gen

.PHONY: all
all: fmt lint gen

PUBLISHED_DATE := $(shell date +%Y-%m-%dT%H:%M:%S)
.PHONY: gen
gen: ## Generates README
	go run main.go
	git add README.md awesome-podcasts.json
	git commit -m "Published $(PUBLISHED_DATE)"
	git push origin master

.PHONY: fmt
fmt: fmt-json fmt-go ## Formats files

.PHONY: lint
lint: ## Lints files
	golangci-lint run

.PHONY: fmt-json
fmt-json: ## Format JSON files
	jq . awesome-podcasts.json > temp.json
	cat temp.json > awesome-podcasts.json
	$(RM) -f temp.json

.PHONY: fmt-go
fmt-go: ## Format Go files
	goimports -w ./
