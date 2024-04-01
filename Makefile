default: help

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: vendor
vendor: ## vendor dependencies
	go mod tidy
	go mod vendor

.PHONY: language-basics-e1
language-basics-e1: ## run language-basics-e1
	go run 02-language-basics/e1/*.go

.PHONY: language-basics-e2
language-basics-e2: ## run language-basics-e2
	go run 02-language-basics/e2/*.go

.PHONY: language-basics-e3
language-basics-e3: ## run language-basics-e3
	cd ./02-language-basics/e3 && go run *.go

.PHONY: language-basics-e4
language-basics-e4: ## run language-basics-e4
	cd ./02-language-basics/e4 && go run *.go

.PHONY: command-line-tools-e1
command-line-tools-e1: ## run command-line-tools-e1
	cd ./03-command-line-tools/e1 && go run *.go proverbs.txt

.PHONY: command-line-tools-e1.1
command-line-tools-e1.1: ## run command-line-tools-e1.1
	cd ./03-command-line-tools/e1.1 && go run *.go proverbs.txt
	cd ./03-command-line-tools/e1.1 && go run *.go fake-text-file.txt

.PHONY: command-line-tools-e2
command-line-tools-e2: ## run command-line-tools-e2
	go run ./03-command-line-tools/e2/*.go -f ./03-command-line-tools/e1/proverbs.txt
	FILE=./03-command-line-tools/e1/proverbs.txt go run ./03-command-line-tools/e2/*.go

.PHONY: command-line-tools-e3
command-line-tools-e3: ## run command-line-tools-e3
	go run ./03-command-line-tools/e3/*.go -f ./03-command-line-tools/e1/proverbs.txt

.PHONY: command-line-tools-e3.1
command-line-tools-e3.1: ## run command-line-tools-e3.1
	go run ./03-command-line-tools/e3.1/*.go -f ./03-command-line-tools/e1/proverbs.txt

.PHONY: command-line-tools-e3.2
command-line-tools-e3.2: ## run command-line-tools-e3.2
	go run ./03-command-line-tools/e3.2/*.go

interfaces-e1: ## run interfaces-e1
	go run ./05-interfaces/e1/*.go

interfaces-e2: ## run interfaces-e2
	go test -v ./05-interfaces/e2/...