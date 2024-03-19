default: help

.PHONY: help
help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

language-basics-e1: ## run language-basics-e1
	@go run 02-language-basics/e1/*.go

language-basics-e2: ## run language-basics-e2
	@go run 02-language-basics/e2/*.go

language-basics-e3: ## run language-basics-e3
	@cd ./02-language-basics/e3 && go run *.go

language-basics-e4: ## run language-basics-e4
	@cd ./02-language-basics/e4 && go run *.go