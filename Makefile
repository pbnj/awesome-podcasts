.DEFAULT_GOAL := gen

.PHONY: fmt
fmt: ## Format Go files
	goimports -w ./

.PHONY: lint
lint: ## Lints files
	golangci-lint run

.PHONY: gen
gen: ## Generates README
	go run main.go

.PHONY: fmt-json
fmt-json: ## Format JSON
	prettier --write awesome-podcasts.json

.PHONY: fmt-md
fmt-md: ## Format markdown
	prettier --write README.md

PUBLISHED_DATE := $(shell date +%Y-%m-%dT%H:%M:%S)
.PHONY: publish
publish: gen fmt-json fmt-md ## Generate, format files, and publish changes
	git add README.md awesome-podcasts.json
	git commit -m "Published $(PUBLISHED_DATE)"
	git push origin master