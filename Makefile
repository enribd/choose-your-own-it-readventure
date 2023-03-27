app ?= it-readventure
version = latest
registry = ghcr.io/enribd
repository = ${registry}/${app}
image = ${repository}:${version}
contents := "index,book-index,author-index,learning-paths,badges,about,mentions"


help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

enable-pre-commit-hook: ### Create git pre-commit hook
	$(info: Create pre-commit file .git/hooks/pre-commit)
	@echo -e "#!/bin/bash\n\nmake pre-commit" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: enable-pre-commit-hook

pre-commit: # run format lint-gh-actions
	@git add README.md
	@git add ./content
	@git add ./mkdocs
.PHONY: pre-commit

lint-gh-actions: ### Lint Github Actions files with actionlint
	$(info Ling Github Actions)
	@docker run --rm \
		-v ${PWD}:/workflows \
		-w /workflows \
		rhysd/actionlint:latest -color
.PHONY: lint-gh-actions

format: ### Run go fmt
	$(info Format code)
	@go fmt ./...
.PHONY: format

run: ### Run go binary (vars: contents)
	$(info Build and run code)
	@mkdir -p ./mkdocs/docs/{more,references,stylesheets}
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --contents=${contents}
.PHONY: run

debug: ### Run go binary in debug mode (vars: contents)
	$(info Build and run code in debug mode)
	@go mod tidy && go mod download && CGO_ENABLED=0 go run main.go --debug --contents=${contents}
.PHONY: debug

mkdocs-build-docker: ## Build Mkdocs docker image
	$(info: Build Mkdocs docker image)
	@docker image build -t ${image} .
.PHONY: build

mkdocs-build-site: ## Build site
	$(info: Build site)
	@docker container run --rm -t --name mkdocs \
		--user $(shell id -u):$(shell id -g) \
		--workdir /src/mkdocs \
		--volume ${PWD}:/src \
		--volume ${PWD}/assets/books/covers:/src/mkdocs/docs/assets/books/covers \
		--volume ${PWD}/assets/learning-paths/icons:/src/mkdocs/docs/assets/learning-paths/icons \
		--volume ${PWD}/assets/favicon.png:/src/mkdocs/docs/assets/favicon.png \
		--volume ${PWD}/assets/logo.png:/src/mkdocs/docs/assets/logo.png \
		${image} build
.PHONY: build

mkdocs-run: ## Run site in localhost
	$(info: Run site in localhost)
	# @xdg-open http://localhost:8000
	@docker container run --rm -t --name mkdocs \
	  --publish 8000:8000 \
		--user $(shell id -u):$(shell id -g) \
		--workdir /src/mkdocs \
		--volume ${PWD}:/src \
		--volume ${PWD}/assets/books/covers:/src/mkdocs/docs/assets/books/covers \
		--volume ${PWD}/assets/learning-paths/icons:/src/mkdocs/docs/assets/learning-paths/icons \
		--volume ${PWD}/assets/favicon.png:/src/mkdocs/docs/assets/favicon.png \
		--volume ${PWD}/assets/logo.png:/src/mkdocs/docs/assets/logo.png \
		${image}
.PHONY: run

clean: ## Clean files
	$(info: Clean files)
	@rm -rf /site
	@rm -rf mkdocs/site
	@rm -rf mkdocs/assets
	@rm -rf mkdocs/docs
.PHONY: clean
