-include .env
-include .env.local
GO ?= go
GOCOVERDIR ?= coverage/integration
export

PACKAGE ?= $(shell ${GO} list -m)

COVERAGE_PACKAGES = ${PACKAGE}

SOURCE_FILES := $(shell ${GO} list -f '{{ range .GoFiles }}{{ printf "%s/%s\n" $$.Dir . }}{{ end }}' ./...)

.PHONY: coverage
coverage: coverage/merged.txt
coverage:
	$(info coverage report)
	@${GO} tool cover -func="$<"

coverage/merged.txt: coverage/unit.txt coverage/integration.txt
	@mkdir -p "$(dir $@)"
	@${GO} run github.com/wadey/gocovmerge@latest $^ > $@

coverage/unit.txt: ${SOURCE_FILES}
coverage/unit.txt:
	@mkdir -p "$(dir $@)"
	$(info running unit tests)
	@${GO} test -v -race -mod=readonly -coverprofile=$@ ./...

coverage/integration.txt: ${GOCOVERDIR}
	@mkdir -p "$(dir $@)"
	$(info generating gcov data)
	@${GO} tool covdata textfmt -i="$<" -o="$@" -pkg="${COVERAGE_PACKAGES}"

${GOCOVERDIR}: ${SOURCE_FILES}
${GOCOVERDIR}:
	$(info running integration test)
	@mkdir -p "$@"
	@${GO} run -cover -race -mod=readonly ./cmd/internal/integration/...
	@${GO} tool covdata percent -i="$@" -pkg="${COVERAGE_PACKAGES}" | column -t
