SHELL   := /bin/bash
VERSION := v1.0.1
GOOS      := $(shell go env GOOS)
GOARCH    := $(shell go env GOARCH)

.PHONY: all
all: vet test build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" ./cmd/qm

.PHONY: package
package: clean build
	gzip qm -c > qm_$(VERSION)_$(GOOS)_$(GOARCH).gz
	sha1sum qm_$(VERSION)_$(GOOS)_$(GOARCH).gz > qm_$(VERSION)_$(GOOS)_$(GOARCH).gz.sha1sum

.PHONY: clean
clean:
	rm -f qm

.PHONY: vet
vet:
	go vet

.PHONY: test
test:
	go test
