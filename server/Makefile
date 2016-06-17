.PHONY: all install test docker

IMAGE ?= colelawr/artsy-host

all: install test

install:
	go get

test:
	@echo "Currently we don't provide test cases!"

docker:
	docker build --rm -t $(IMAGE) .
