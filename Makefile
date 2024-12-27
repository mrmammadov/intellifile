.PHONY: build install

build:
	go build -o bin/intellifile cmd/intellifile/main.go

install:
	go install ./cmd/intellifile