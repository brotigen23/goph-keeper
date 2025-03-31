BUILDDIR=bin

all:
	cd client/
	make

build:
	go build -o ${BUILDDIR}/client cmd/client/main.go

create-docker-images:


docker-image-client:
	docker build --tag goph-keeper-client ./client


.PHONY: all build all docker-image
