help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

enable-pre-commit-hook: ### Create git pre-commit hook
	$(info: Create pre-commit file .git/hooks/pre-commit)
	@echo -e "#!/bin/bash\n\nmake pre-commit" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
.PHONY: enable-pre-commit-hook

pre-commit: resize-images
.PHONY: pre-commit

resize-images: ## Resize all images
	$(info: Resize all images)
	@find ./assets/covers -type f -exec convert {} -resize 122x160 {} \;
	@git add ./assets/covers/
.PHONY: resize-images
