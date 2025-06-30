GOLANGCI_LINT_VERSION := 2.1.6

GIT_VERSION := $(shell git describe --tags --always)
GIT_REVISION := $(shell git rev-parse HEAD)

GO_BUILD_LDFLAGS := "-X github.com/kukymbr/sqlamble/internal/version.Version=$(GIT_VERSION) \
                -X github.com/kukymbr/sqlamble/internal/version.Revision=$(GIT_REVISION) \
                -X github.com/kukymbr/sqlamble/internal/version.BuiltAt=$(shell date -u +%Y%m%d%H%M%S)"
GO_BUILD_ARGS := $(build_arguments) --ldflags $(GO_BUILD_LDFLAGS)

all:
	$(MAKE) clean
	$(MAKE) prepare
	$(MAKE) validate
	$(MAKE) build

prepare:
	go install ./...
	go fmt ./...
	$(MAKE) generate

validate:
	go vet ./...
	$(MAKE) lint
	$(MAKE) test

build:
	go build $(GO_BUILD_ARGS) -o bin/cleemebackend ./cmd/backend

generate:
	go generate ./...
	go mod tidy


lint:
	if [ ! -f ./bin/golangci-lint ]; then \
  		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "./bin" "v${GOLANGCI_LINT_VERSION}"; \
  	fi;
	./bin/golangci-lint run ./...

test:
	CGO_ENABLED=1 go test -race -coverprofile=coverage_out ./...
	go tool cover -func=coverage_out
	go tool cover -html=coverage_out -o coverage.html
	rm -f coverage_out

test_report:
	CGO_ENABLED=1 go test -race -coverprofile=coverage_out -v 2>&1 ./... | go-junit-report -set-exit-code -iocopy -out junit.report.xml
	go tool cover -func=coverage_out
	go tool cover -html=coverage_out -o coverage.html

test_short:
	go test -short ./...

clean:
	go clean
