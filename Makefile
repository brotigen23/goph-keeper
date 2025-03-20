BUILDDIR=bin

all:
	go run cmd/client/main.go

build: 
	go build -o ${BUILDDIR}/client cmd/client/main.go




.PHONY: build all
