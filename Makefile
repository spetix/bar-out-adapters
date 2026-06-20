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
	mkdir -p coverage
	go tool cover -html=coverage.txt -o coverage/index.html

coverage.xml:
	gocover-cobertura < coverage.txt > coverage.xml

project-stats: test-results.json
	python3 scripts/generate-project-stats.py

prepare-site:
	mkdir -p build/site
	cp README.md build/site/
	cp bin/* build/site
	cp code-coverage-results.md build/site
	cp coverage.xml build/site

inject-coverage:
	mkdir -p $(DOCS_DIR)/static/coverage
	if [ -d coverage ]; then \
		cp -r coverage/* $(DOCS_DIR)/static/coverage/; \
	else \
		echo "warning: coverage directory not found, skipping coverage injection"; \
	fi

inject-test-results:
	cp test-results.json $(DOCS_DIR)/static/test-results.json

docs-setup:
	npm --prefix $(DOCS_DIR) ci

docs-build:
	npm --prefix $(DOCS_DIR) run build

docs-prepare: project-stats inject-coverage inject-test-results

docs-site: docs-prepare docs-build

local-site: docs-setup coverage docs-site
	rm -rf $(DOCS_DIR)/build
	podman run --rm -v "$(PWD)/$(DOCS_DIR):/site" -w /site node:25 npm install
	npm --prefix $(DOCS_DIR) run build

run-local-site: local-site
	npm --prefix $(DOCS_DIR) run serve

run: docs-dev

update-versions:
	python3 scripts/update-versions.py

docs-dev:
	npm --prefix $(DOCS_DIR) run start

docs-serve:
	npm --prefix $(DOCS_DIR) run serve

docs-clean:
	rm -rf $(DOCS_DIR)/build $(DOCS_DIR)/.docusaurus $(DOCS_DIR)/node_modules

all: build test coverage project-stats
