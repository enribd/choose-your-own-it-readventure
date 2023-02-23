contents := "readme,book-index,author-index,learning-paths"

help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

enable-pre-commit-hook: ### Create git pre-commit hook
	$(info: Create pre-commit file .git/hooks/pre-commit)
	@echo -e "#!/bin/bash\n\nmake pre-commit" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: enable-pre-commit-hook

pre-commit: run format resize-images #lint-gh-actions
	@git add ./content
.PHONY: pre-commit

resize-images: ## Resize all images
	$(info: Resize all images)
	@find ./assets/books/covers -type f -exec convert {} -resize 122x160 {} \;
	@git add ./assets/books/covers
.PHONY: resize-images

lint-gh-actions: ### Lint Github Actions files with actionlint
	$(info Ling Github Actions)
	@docker run --rm \
		-v $(shell pwd):/workflows \
		-w /workflows \
		rhysd/actionlint:latest -color
.PHONY: lint-gh-actions

format: ### Run go fmt
	$(info Format code)
	@go fmt ./...
.PHONY: format

run: ### Run go binary (vars: contents)
	$(info Build and run code)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --contents=${contents}
.PHONY: run

debug: ### Run go binary in debug mode (vars: contents)
	$(info Build and run code in debug mode)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --debug --contents=${contents}
.PHONY: debug

trace: ### Run go binary in trace mode (vars: contents)
	$(info Build and run code in trace mode)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --debug --trace --contents=${contents}
.PHONY: trace
