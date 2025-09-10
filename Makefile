ARGS?=
PROJECT_NAME?=bar-out-adapters

setup:
	go install github.com/boumenot/gocover-cobertura@latest || exit 0
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest || exit 0
	go mod download
	go mod tidy

lib/$(PROJECT_NAME)-linux-amd64: 
	@echo "Building linux-amd64"
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -v -o lib/$(PROJECT_NAME)-linux-amd64 -a ./pkg/barout/blocketOutputInterface.go



lib/$(PROJECT_NAME)-linux-arm64:
	@echo "Building linux-arm64"
	mkdir -p bin
	GOOS=linux GOARCH=arm64 go build -v -o lib/$(PROJECT_NAME)-linux-arm64 -a ./pkg/barout/blocketOutputInterface.go



# Define the build-all target
.PHONY: build
build: setup
	$(MAKE) lib/$(PROJECT_NAME)-linux-amd64
	$(MAKE) lib/$(PROJECT_NAME)-linux-arm64


# TODO clean binaries
clean:
	go clean ./pkg/...
	rm -Rf lib
	rm -f test-result*.json
	rm -f coverage.*

test: build
	$(MAKE) test-results.json

# TODO run tests
test-results.json:
	go test -race -json -v -coverprofile=coverage.txt ./... 2>&1 | tee test-results.json | gotestfmt

coverage: test
	$(MAKE) coverage.xml

coverage.xml:
	gocover-cobertura < coverage.txt > coverage.xml


prepare-site:
	mkdir -p build/site
	cp README.md build/site/
	cp bin/* build/site
	cp code-coverage-results.md build/site
	cp coverage.xml build/site

all: build test coverage
