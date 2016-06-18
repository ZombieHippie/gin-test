.PHONY: all clean deps fmt vet test docker

EXECUTABLE = dryclean-host
IMAGE = colelawr/$(EXECUTABLE)
COMMIT ?= $(shell git rev-parse --short HEAD)

LDFLAGS = -X "main.buildCommit=$(COMMIT)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: deps build test

clean:
	go clean -i ./...

deps:
	go get

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

docker:
	GOOS=linux go build  -a -installsuffix cgo -ldflags '-s -w $(LDFLAGS)'
	docker build --rm -t $(IMAGE) .

$(EXECUTABLE): $(wildcard main.go)
	go build -ldflags '-s -w $(LDFLAGS)'

build: $(EXECUTABLE)
