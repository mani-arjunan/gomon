OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
PREFIX ?= $(HOME)/.local
BINDIR := $(PREFIX)/bin

init:
	go mod tidy

clean:
	rm -rf bin gomon

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o gomon

install: clean build
	mkdir -p $(BINDIR)
	mv gomon $(BINDIR)/
	@echo "Installed binaries to $(BINDIR)"

