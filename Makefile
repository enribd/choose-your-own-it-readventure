help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

enable-pre-commit-hook: ### Create git pre-commit hook
	$(info: Create pre-commit file .git/hooks/pre-commit)
	@echo -e "#!/bin/bash\n\nmake pre-commit" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: enable-pre-commit-hook

pre-commit: format-code clean resize-images #lint-gh-actions
.PHONY: pre-commit

resize-images: ## Resize all images
	$(info: Resize all images)
	@find ./assets/covers -type f -exec convert {} -resize 122x160 {} \;
	@git add ./assets/covers/
.PHONY: resize-images

lint-gh-actions: ### Lint Github Actions files with actionlint
	$(info Ling Github Actions)
	@docker run --rm \
		-v $(shell pwd):/workflows \
		-w /workflows \
		rhysd/actionlint:latest -color
.PHONY: lint-gh-actions

format-code: ### Run go fmt
	$(info Format code)
	@go fmt ./...
.PHONY: format-code

run: ### Run go binary
	$(info Build and run code)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go
.PHONY: run

debug: ### Run go binary in debug mode
	$(info Build and run code in debug mode)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --debug
.PHONY: debug

clean: ## Clean unneeded files
	$(info: Clean unneeded files)
	@find . -iname '*-test*.md' -exec rm {} \;
.PHONY: clean
