GOCOVERDIR ?= build/coverage/integration
GOFLAGS += -covermode=atomic -race -shuffle=on -mod=readonly -trimpath

export GOCOVERDIR GOFLAGS

GOMODULE != go list -m

define GO_TEST_IGNORE_PATTERNS
$(strip
cmd/internal/integration/main.go:
)
endef

define TEST_PACKAGES
$(strip
)
endef

GO_SOURCES = $(shell git ls-files '*.go')

all:: gomod2nix.toml

check:: GOFLAGS=
check::
	@golangci-lint run

.PHONY: coverage
#: Reports code coverage.
coverage: build/coverage/all.txt | build/coverage/
	@go tool cover -func=$< | sed 's|${GOMODULE}/||g' | column -t

.PHONY: uncovered
#: Reports uncovered/partially covered code.
uncovered: build/coverage/all.txt | build/coverage/
	@go tool cover -func=$< | sed 's|${GOMODULE}/||g' | column -t | grep -v '100.0%' || true

test::
	@gotestdox ./...
	@go run -cover ./cmd/internal/integration/...

.PHONY: coverage-html
#: Generates an HTML coverage report.
coverage-html: build/coverage/all.html

## make magic, not war ;)

build/coverage/all.txt: build/coverage/unit.part.txt build/coverage/integration.part.txt
	go tool gocovmerge $^ \
	| grep $(if ${GO_TEST_IGNORE_PATTERNS},-v '$(subst $(space),\|,${GO_TEST_IGNORE_PATTERNS})',.) \
	> $@

build/coverage/unit.part.txt: ${GO_SOURCES} go.sum | build/coverage/
	gotestdox --coverprofile=$@ $(addprefix ./,$(addsuffix /...,${TEST_PACKAGES}))
	sed -i'' '#$(subst .,\.,$(subst $(space),\|,${GO_TEST_IGNORE_PATTERNS}))#d' $@

build/coverage/integration.part.txt: ${GO_SOURCES} go.sum | ${GOCOVERDIR}/
	-@rm -rf "${GOCOVERDIR}/*" 2>/dev/null || true
	go run -cover ./cmd/internal/integration/...
	go tool covdata textfmt -i=${GOCOVERDIR} -o=$@ $(if ${TEST_PACKAGES},-pkg="$(addprefix ${GOMODULE}/,${TEST_PACKAGES})")
	sed -i'' '#$(subst .,\.,$(subst $(space),\|,${GO_TEST_IGNORE_PATTERNS}))#d' $@

bin/%: ${GO_SOURCES} go.sum
	go build -o ./$@ ./cmd/$(patsubst bin/%,%,$@)/...

go.sum: GOFLAGS-=-mod-readonly
go.sum: ${GO_SOURCES} go.mod
	@go mod tidy -v -x
	@touch $@

gomod2nix.toml: go.sum
	gomod2nix generate

build/coverage/%.html: build/coverage/%.txt | build/coverage/
	go tool cover -html=$< -o $@
