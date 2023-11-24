VERSION := $(shell git describe --tags)

GOLDFLAGS += -X main.Version=$(VERSION) -H=windowsgui
GOFLAGS = -ldflags "$(GOLDFLAGS)"

run: build
	./build/cliptr.exe

build:
	go-winres make --product-version=git-tag
	go build -o ./build/cliptr.exe $(GOFLAGS) 
	cp README.md ./build
	cp LICENSE ./build
