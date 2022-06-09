.DEFAULT_GOAL := help
SHELL = /bin/bash

GIT_HOOKS = .git/hooks/commit-msg .git/hooks/pre-commit .git/hooks/pre-push .git/hooks/prepare-commit-msg

$(GIT_HOOKS): .git/hooks/%: .githooks/%

.githooks/%:
	touch $@

.git/hooks/%:
	cp $< $@

.PHONY: add-git-configs
add-git-configs: ## Add Git Configs
	git config --global branch.autosetuprebase always
	git config --global color.branch true
	git config --global color.diff true
	git config --global color.interactive true
	git config --global color.status true
	git config --global color.ui true
	git config --global commit.gpgsign true
	git config --global core.autocrlf input
	git config --global core.editor "code --wait"
	git config --global diff.tool code
	git config --global difftool.code.cmd "code --diff \$$LOCAL \$$REMOTE --wait"
	git config --global gpg.program gpg
	git config --global init.defaultbranch main
	git config --global log.date relative
	git config --global merge.tool code
	git config --global mergetool.code.cmd "code --wait \$$MERGED"
	git config --global pull.default current
	git config --global pull.rebase true
	git config --global push.default current
	git config --global rebase.autostash true
	git config --global rerere.enabled true
	git config --global stash.showpatch true
	git config --global tag.gpgsign true

.PHONY: add-git-hooks
add-git-hooks: clean-git-hooks $(GIT_HOOKS) ## Add Git Hooks

.PHONY: audit
audit: ## Audit
	echo "audit"

.PHONY: check
check: ## Check
	echo "check"

.PHONY: clean-git-hooks
clean-git-hooks: ## Clean Git Hooks
	rm -fr $(GIT_HOOKS)

.PHONY: clippy
clippy: ## Clippy
	echo "clippy"

.PHONY: conventional-commits-linter
conventional-commits-linter: ## Conventional Commits Linter
	echo "conventional-commits-linter"

.PHONY: coverage
coverage: ## Coverage
	go test ./... -covermode=atomic -coverprofile=./coverage/coverage.out
	go tool cover -html=./coverage/coverage.out -o=./coverage/coverage.html

.PHONY: deny-check
deny-check: ## Deny Check
	echo "deny-check"

.PHONY: fmt-check
fmt-check: ## FMT Check
	go fmt .

.PHONY: git
git: add-git-configs add-git-hooks ## Add Git Configs & Hooks

.PHONY: help
help: ## Help
	@grep --extended-regexp '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m%-33s\033[0m %s\n", $$1, $$2 }'

.PHONY: test
test: ## Test
	go test
