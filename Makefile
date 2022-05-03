.PHONY: help

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display this help message.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary of devctl
	rm -rf dist || true
	mkdir dist
	cd src && env GOOS=linux GOARCH=amd64 go build -v -o devctl
	mv src/devctl dist/devctl
