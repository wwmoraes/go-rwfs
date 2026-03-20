GOFLAGS += -covermode=atomic -race -shuffle=on -mod=readonly -trimpath

export GOFLAGS

GOMODULE != go list -m
GO_SOURCES != git ls-files '*.go'

all:: gomod2nix.toml

check:: GOFLAGS=
check::
	golangci-lint run

clean::
	rm -rf build

.PHONY: coverage
#: Reports code coverage.
coverage: build/coverage/unit.txt | build/coverage/
	go tool cover -func=$< | sed 's|${GOMODULE}/||g' | column -t

.PHONY: uncovered
#: Reports uncovered/partially covered code.
uncovered: build/coverage/unit.txt | build/coverage/
	go tool cover -func=$< | sed 's|${GOMODULE}/||g' | column -t | grep -v '100.0%' || true

test::
	go test -v ./...

.PHONY: coverage-html
#: Generates an HTML coverage report.
coverage-html: build/coverage/unit.html
	open $<

## make magic, not war ;)

build/coverage/unit.txt: ${GO_SOURCES} go.sum | build/coverage/
	go test -v --coverprofile=$@ ./...

go.sum: GOFLAGS-=-mod-readonly
go.sum: ${GO_SOURCES} go.mod
	go mod tidy -v -x
	@touch $@

gomod2nix.toml: go.sum
	gomod2nix generate

build/coverage/%.html: build/coverage/%.txt | build/coverage/
	go tool cover -html=$< -o $@
