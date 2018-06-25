.PHONY: all
all: fmt gen

.PHONY: gen
gen: ## Generates README
	vgo run main.go -gen

.PHONY: fmt
fmt: ## Formats JSON file
	vgo run main.go -fmt
