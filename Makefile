.PHONY: all
all: fmt lint

.PHONY: gen
gen: ## Generates README
	@go run main.go -gen

.PHONY: fmt
fmt: fmt-json fmt-go ## Formats files

.PHONY: lint
lint: ## Lints files
	@golangci-lint run

.PHONY: fmt-json
fmt-json: ## Format JSON files
	@jq . awesome-podcasts.json > temp.json
	@cat temp.json > awesome-podcasts.json
	@$(RM) -f temp.json

.PHONY: fmt-go
fmt-go: ## Format Go files
	@goimports -w ./
