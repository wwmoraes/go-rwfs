-include .env
-include .env.local
export

MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
MAKEFLAGS += --jobs --max-load

SHELL := $(shell which bash)
.SHELLFLAGS := -euo pipefail -c
.DEFAULT_GOAL := all

.PHONY: all
#: Builds the entire project.
all:: ;

.PHONY: check
#: Perform self-checks such as linting and formatting.
check::
	nix flake check

.PHONY: clean
#: Delete all files that are normally created by building.
clean:: ;

.PHONY: test
#: Builds and runs tests.
test:: ;

-include $(shell git ls-files '**/*.mk')
