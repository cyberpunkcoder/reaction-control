.DEFAULT_GOAL := default

.PHONY: default
default: build

.PHONY: cook-rice
cook-rice:
	cd ./cmd/reaction-control && rice embed-go

.PHONY: build
build: cook-rice
	cd ./cmd/reaction-control && mkdir -p ../../build && go build -o ../../build/reaction-control