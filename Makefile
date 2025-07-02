GOLANGCI_LINT_VERSION := 2.1.6

GIT_VERSION := $(shell git describe --tags --always)
GIT_REVISION := $(shell git rev-parse HEAD)
VERSION_PACKAGE := "github.com/kukymbr/sqlamble/internal/version"

GO_BUILD_LDFLAGS := "-X $(VERSION_PACKAGE).Version=$(GIT_VERSION) \
                -X $(VERSION_PACKAGE).Revision=$(GIT_REVISION) \
                -X $(VERSION_PACKAGE).BuiltAt=$(shell date -u +%Y%m%d%H%M%S)"
GO_BUILD_ARGS := $(build_arguments) --ldflags $(GO_BUILD_LDFLAGS)

all:
	$(MAKE) clean
	$(MAKE) prepare
	$(MAKE) validate
	$(MAKE) build

prepare:
	go install ./...
	go fmt ./...

validate:
	go vet ./...
	$(MAKE) lint
	$(MAKE) test

build:
	go build $(GO_BUILD_ARGS) -o bin/sqlamble ./cmd/sqlamble

lint:
	if [ ! -f ./bin/golangci-lint ]; then \
  		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "./bin" "v${GOLANGCI_LINT_VERSION}"; \
  	fi;
	./bin/golangci-lint run ./...

test:
	go test -coverprofile=coverage_out ./...
	go tool cover -func=coverage_out
	go tool cover -html=coverage_out -o coverage.html
	rm -f coverage_out

test_report:
	go test -coverprofile=coverage_out -v 2>&1 ./... | go-junit-report -set-exit-code -iocopy -out junit.report.xml
	go tool cover -func=coverage_out
	go tool cover -html=coverage_out -o coverage.html

test_short:
	go test -short ./...

clean:
	go clean

generate_example:
	go build \
		--ldflags "-X $(VERSION_PACKAGE).BuiltAt=$(shell date -u +%Y%m%d%H%M%S)" \
		-o bin/sqlamble_example \
		./cmd/sqlamble
	./bin/sqlamble_example --source=example/sql --target=example/internal/queries --fmt=gofmt
	rm ./bin/sqlamble_example