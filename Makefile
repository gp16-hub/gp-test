
# 定义api目标
.PHONY: check-branch

.DEFAULT_GOAL := check-branch

# Colors
RED=\033[0;31m
GREEN=\033[0;32m
NC=\033[0m

CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Define a target to check if the current branch is master
check-branch:
	@if [ "$(CURRENT_BRANCH)" = "master" ]; then \
		echo "${GREEN}protocol are on the master branch.${NC}"; \
	else \
		echo  "${RED}protocol are not on the master branch. protocol are on the $(CURRENT_BRANCH) branch.${NC}"; \
	fi


.PHONY:proto-format
proto-format:
  # brew install clang-format
	@./scripts/proto-format.sh

.PHONY:gen
gen:
	@./scripts/gen.sh
