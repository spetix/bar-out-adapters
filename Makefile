ARGS?=
PROJECT_NAME?=bar-out-adapters
DOCS_DIR=docs-site

setup:
	go install github.com/boumenot/gocover-cobertura@latest || exit 0
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest || exit 0
	go install github.com/vektra/mockery/v3@latest || exit 0
	go mod download

lib/$(PROJECT_NAME)-linux-amd64: 
	@echo "Building linux-amd64"
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -v -o lib/$(PROJECT_NAME)-linux-amd64 -a ./pkg/barout/api.go



lib/$(PROJECT_NAME)-linux-arm64:
	@echo "Building linux-arm64"
	mkdir -p bin
	GOOS=linux GOARCH=arm64 go build -v -o lib/$(PROJECT_NAME)-linux-arm64 -a ./pkg/barout/api.go



# Define the build-all target
.PHONY: build
build: mockery
	go mod tidy
	$(MAKE) lib/$(PROJECT_NAME)-linux-amd64
	$(MAKE) lib/$(PROJECT_NAME)-linux-arm64


# TODO clean binaries
clean:
	go clean ./pkg/...
	rm -Rf lib
	rm -f test-result*.json
	rm -f coverage.*
	rm -Rf generated

mockery: setup
	mockery --config .mockery.yml

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


docs-dev:
	cd $(DOCS_DIR) && npm run start

docs-build:
	cd $(DOCS_DIR) && npm run build

docs-serve:
	cd $(DOCS_DIR) && npm run serve

docs-clean:
	cd $(DOCS_DIR) && rm -rf build .docusaurus node_modules

all: build test coverage
